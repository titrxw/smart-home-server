package logic

import (
	"context"
	"reflect"
	"time"

	exception "github.com/titrxw/smart-home-server/app/Exception"
	"gorm.io/gorm"

	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
)

type DeviceOperateLogic struct {
	LogicAbstract
}

func (deviceOperateLogic DeviceOperateLogic) IsSuccessResponse(operatePayload model.OperatePayload) (bool, error) {
	if _, ok := operatePayload["status"]; !ok {
		return false, exception.NewArgsError("status 参数缺失")
	}
	if _, ok := operatePayload["status"].(string); !ok {
		return false, exception.NewArgsError("status 参数格式错误")
	}

	if operatePayload["status"] == "success" {
		return true, nil
	}

	return false, nil
}

func (deviceOperateLogic DeviceOperateLogic) TriggerOperate(ctx context.Context, device *model.Device, operate model.DeviceOperateType, payload model.OperatePayload, operateLevel uint8) (*model.DeviceOperateLog, error) {
	if device.IsDelete() {
		return nil, exception.NewLogicError("该设备已删除")
	}
	if device.IsDisable() {
		return nil, exception.NewLogicError("该设备状态异常")
	}
	if !Logic.DeviceLogic.IsSupportOperate(device, operate) {
		return nil, exception.NewLogicError("该设备不支持当前操作")
	}
	if !device.IsOnline() {
		return nil, exception.NewLogicError("该设备当前不在线")
	}

	gatewayDevice := device
	if Logic.DeviceLogic.IsNeedGateway(device) {
		gatewayDevice = Logic.DeviceGatewayLogic.GetGatewayDevice(ctx, device)
		if gatewayDevice == nil {
			return nil, exception.NewLogicError("该设备未绑定网关")
		}
	}

	deviceOperateLog := &model.DeviceOperateLog{
		DeviceId:        device.ID,
		DeviceGatewayId: gatewayDevice.ID,
		DeviceType:      device.TypeName,
		Source:          global.FApp.Name,
		OperateName:     string(operate),
		OperateNumber:   deviceOperateLogic.GetOperateOrReportNumber(device.App.AppId),
		OperateTime:     model.LocalTime(time.Now()),
		OperatePayload:  payload,
		OperateLevel:    operateLevel,
		CreatedAt:       model.LocalTime(time.Now()),
	}

	gatewayAdapter := Logic.DeviceLogic.GetDeviceAdapter(gatewayDevice.TypeName)
	err := gatewayAdapter.BeforeTriggerOperate(ctx, gatewayDevice, device, deviceOperateLog)
	if err != nil {
		return nil, err
	}

	if !repository.Repository.DeviceOperateLogRepository.AddDeviceOperateLog(deviceOperateLogic.GetDefaultDb(), deviceOperateLog) {
		return nil, exception.NewLogicError("添加操作记录失败")
	}

	message := &model.IotMessage{
		EventType: deviceOperateLog.OperateName,
		Id:        deviceOperateLog.OperateNumber,
		Payload:   deviceOperateLog.OperatePayload,
		Timestamp: time.Time(deviceOperateLog.CreatedAt).Unix(),
	}
	err = Logic.MessageLogic.PubClientOperate(ctx, gatewayDevice, device, message)
	if err != nil {
		deviceOperateLog.ResponsePayload = model.OperatePayload{"error": err.Error()}
		if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(deviceOperateLogic.GetDefaultDb(), deviceOperateLog) {
			return nil, exception.NewLogicError("更新操作记录失败")
		}
	}

	err = gatewayAdapter.AfterTriggerOperate(ctx, gatewayDevice, device, deviceOperateLog)
	if err != nil {
		return nil, err
	}

	//触发事件
	global.FApp.Event.Publish(reflect.TypeOf(event.DeviceOperateTriggerEvent{}).Name(), event.NewDeviceOperateTriggerEvent(device, deviceOperateLog))

	return deviceOperateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) OnOperateResponse(gatewayDevice *model.Device, device *model.Device, iotMessage *model.IotMessage) error {
	operateLog, err := deviceOperateLogic.GetOperateLogByNumber(iotMessage.Id)
	if err == nil {
		if device.ID != operateLog.DeviceId {
			return exception.NewLogicError("设备不匹配")
		}
		err = deviceOperateLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
			operateLog.ResponsePayload = iotMessage.Payload
			operateLog.ResponseTime = time.Unix(iotMessage.Timestamp, 0).Format(model.TimeFormat)
			if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(tx, operateLog) {
				return exception.NewLogicError("更新操作记录失败")
			}

			status, exists := operateLog.ResponsePayload["cur_status"]
			if exists && reflect.TypeOf(status).Name() == "string" {
				device.DeviceCurStatus = status.(string)
				if !repository.Repository.DeviceRepository.UpdateDevice(tx, device) {
					return exception.NewLogicError("更新设备失败")
				}
			}

			return nil
		})

		if err == nil {
			err = Logic.DeviceLogic.GetDeviceAdapter(gatewayDevice.TypeName).OnOperateResponse(context.Background(), gatewayDevice, device, operateLog, iotMessage)
			global.FApp.Event.Publish(reflect.TypeOf(event.DeviceOperateReplyEvent{}).Name(), event.NewDeviceOperateReplyEvent(device, operateLog, iotMessage))
		}
	}

	return err
}

func (deviceOperateLogic DeviceOperateLogic) UpdateOperateLog(operateLog *model.DeviceOperateLog) error {
	if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(deviceOperateLogic.GetDefaultDb(), operateLog) {
		return exception.NewLogicError("更新操作记录失败")
	}

	return nil
}

func (deviceOperateLogic DeviceOperateLogic) GetDeviceOperateLogByNumber(device *model.Device, operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog, _ := deviceOperateLogic.GetOperateLogByNumber(operateNumber)
	if operateLog.DeviceId != device.ID {
		return nil, exception.NewLogicError("非法操作")
	}

	return operateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) GetOperateLogByNumber(operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog := repository.Repository.DeviceOperateLogRepository.GetDeviceOperateLogByOperateNumber(deviceOperateLogic.GetDefaultDb(), operateNumber)
	if operateLog == nil {
		return nil, exception.NewLogicError("设备操作记录不存在")
	}

	return operateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) GetDeviceOperates(device *model.Device, page uint, pageSize uint) *repository.PageModel {
	return repository.Repository.DeviceOperateLogRepository.GetDeviceOperates(deviceOperateLogic.GetDefaultDb(), device.ID, page, pageSize)
}

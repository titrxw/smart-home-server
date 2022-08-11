package logic

import (
	"context"
	"encoding/json"
	"errors"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"gorm.io/gorm"
	"reflect"
	"time"

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
		return false, errors.New("status 参数缺失")
	}
	if _, ok := operatePayload["status"].(string); !ok {
		return false, errors.New("status 参数格式错误")
	}

	if operatePayload["status"] == "success" {
		return true, nil
	}

	return false, nil
}

func (deviceOperateLogic DeviceOperateLogic) TriggerOperate(ctx context.Context, device *model.Device, operate model.DeviceOperateType, payload model.OperatePayload, operateLevel uint8) (*model.DeviceOperateLog, error) {
	if device.IsDelete() {
		return nil, errors.New("该设备已删除")
	}
	if device.IsDisable() {
		return nil, errors.New("该设备状态异常")
	}
	if !Logic.DeviceLogic.IsSupportOperate(device, operate) {
		return nil, errors.New("该设备不支持当前操作")
	}
	if !device.IsOnline() {
		return nil, errors.New("该设备当前不在线")
	}

	deviceOperateLog := &model.DeviceOperateLog{
		DeviceId:       device.ID,
		DeviceType:     device.Type,
		Source:         global.FApp.Name,
		OperateName:    string(operate),
		OperateNumber:  deviceOperateLogic.GetOperateOrReportNumber(device.App.AppId),
		OperateTime:    model.LocalTime(time.Now()),
		OperatePayload: payload,
		OperateLevel:   operateLevel,
		CreatedAt:      model.LocalTime(time.Now()),
	}

	err := Logic.DeviceLogic.GetDeviceAdapter(device.Type).BeforeTriggerOperate(device, deviceOperateLog)
	if err != nil {
		return nil, err
	}

	if !repository.Repository.DeviceOperateLogRepository.AddDeviceOperateLog(deviceOperateLogic.GetDefaultDb(), deviceOperateLog) {
		return nil, errors.New("添加操作记录失败")
	}

	err = Logic.EmqxLogic.PubClientOperate(ctx, device, deviceOperateLog)
	if err != nil {
		deviceOperateLog.ResponsePayload = model.OperatePayload{"error": err.Error()}
		if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(deviceOperateLogic.GetDefaultDb(), deviceOperateLog) {
			return nil, errors.New("更新操作记录失败")
		}
	}

	err = Logic.DeviceLogic.GetDeviceAdapter(device.Type).AfterTriggerOperate(device, deviceOperateLog)
	if err != nil {
		return nil, err
	}

	//触发事件
	global.FApp.Event.Publish(reflect.TypeOf(event.DeviceOperateTriggerEvent{}).Name(), event.NewDeviceOperateTriggerEvent(device, deviceOperateLog))

	return deviceOperateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) OnOperateResponse(device *model.Device, cloudEvent *cloudevents.Event) (*model.DeviceOperateLog, error) {
	payLoad := model.OperatePayload{}
	err := json.Unmarshal(cloudEvent.Data(), &payLoad)
	if err != nil {
		return nil, err
	}

	operateLog, err := deviceOperateLogic.GetOperateLogByNumber(cloudEvent.ID())
	if err == nil {
		if device.ID != operateLog.DeviceId {
			return nil, errors.New("设备不匹配")
		}
		err = deviceOperateLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
			operateLog.ResponsePayload = payLoad
			operateLog.ResponseTime = cloudEvent.Time().Format(model.TimeFormat)
			if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(tx, operateLog) {
				return errors.New("更新操作记录失败")
			}

			status, exists := operateLog.ResponsePayload["cur_status"]
			if exists && reflect.TypeOf(status).Name() == "string" {
				device.DeviceCurStatus = status.(string)
				if !repository.Repository.DeviceRepository.UpdateDevice(tx, device) {
					return errors.New("更新设备失败")
				}
			}

			return nil
		})
	}

	return operateLog, err
}

func (deviceOperateLogic DeviceOperateLogic) UpdateOperateLog(operateLog *model.DeviceOperateLog) error {
	if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(deviceOperateLogic.GetDefaultDb(), operateLog) {
		return errors.New("更新操作记录失败")
	}

	return nil
}

func (deviceOperateLogic DeviceOperateLogic) GetDeviceOperateLogByNumber(device *model.Device, operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog, _ := deviceOperateLogic.GetOperateLogByNumber(operateNumber)
	if operateLog.DeviceId != device.ID {
		return nil, errors.New("非法操作")
	}

	return operateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) GetOperateLogByNumber(operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog := repository.Repository.DeviceOperateLogRepository.GetDeviceOperateLogByOperateNumber(deviceOperateLogic.GetDefaultDb(), operateNumber)
	if operateLog == nil {
		return nil, errors.New("设备操作记录不存在")
	}

	return operateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) GetDeviceOperates(device *model.Device, page uint, pageSize uint) *repository.PageModel {
	return repository.Repository.DeviceOperateLogRepository.GetDeviceOperates(deviceOperateLogic.GetDefaultDb(), device.ID, page, pageSize)
}

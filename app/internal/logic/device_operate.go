package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	"github.com/titrxw/smart-home-server/app/internal/event"
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/internal/repository"
	pkgdevice "github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
	pkgmodel "github.com/titrxw/smart-home-server/app/pkg/model"
	pkgrepository "github.com/titrxw/smart-home-server/app/pkg/repository"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type DeviceOperate struct {
	logic.Abstract
}

func (l DeviceOperate) TriggerOperate(ctx context.Context, device *model.Device, operate model.DeviceOperateType, payload model.OperatePayload, operateLevel uint8) (*model.DeviceOperateLog, error) {
	if device.IsDelete() {
		return nil, exception.NewResponseError("该设备已删除")
	}
	if device.IsDisable() {
		return nil, exception.NewResponseError("该设备状态异常")
	}
	if !Logic.Device.IsSupportOperate(device, operate) {
		return nil, exception.NewResponseError("该设备不支持当前操作")
	}
	if !device.IsOnline() {
		return nil, exception.NewResponseError("该设备当前不在线")
	}

	gatewayDevice := device
	if manager.GetDeviceConfigByDeviceType(device.TypeName).NeedGateway {
		gatewayDevice = Logic.DeviceGateway.GetGatewayDevice(ctx, device)
		if gatewayDevice == nil {
			return nil, exception.NewResponseError("该设备未绑定网关")
		}
	}

	deviceOperateLog := &model.DeviceOperateLog{
		DeviceId:        device.ID,
		DeviceGatewayId: gatewayDevice.ID,
		DeviceType:      device.TypeName,
		Source:          helper.GetAppName(),
		OperateName:     string(operate),
		OperateNumber:   l.GetOperateOrReportNumber(device.App.AppId),
		OperateTime:     pkgmodel.LocalTime(time.Now()),
		OperatePayload:  payload,
		OperateLevel:    operateLevel,
		CreatedAt:       pkgmodel.LocalTime(time.Now()),
	}

	message := &pkgdevice.OperateMessage{
		EventType: deviceOperateLog.OperateName,
		Id:        deviceOperateLog.OperateNumber,
		Payload:   deviceOperateLog.OperatePayload,
		Timestamp: time.Time(deviceOperateLog.CreatedAt).Unix(),
	}

	if !repository.Repository.DeviceOperateLog.AddDeviceOperateLog(l.GetDefaultDb(), deviceOperateLog) {
		return nil, exception.NewResponseError("添加操作记录失败")
	}

	err := Logic.Message.PubClientOperateMsg(ctx, gatewayDevice, device, message)
	if err != nil {
		deviceOperateLog.ResponsePayload = model.OperatePayload{"error": err.Error()}
		if !repository.Repository.DeviceOperateLog.UpdateDeviceOperateLog(l.GetDefaultDb(), deviceOperateLog) {
			return nil, exception.NewResponseError("更新操作记录失败")
		}
	}

	//触发事件
	app.GApp.GetEvent().Publish(reflect.TypeOf(event.DeviceOperateTrigger{}).Name(), event.NewDeviceOperateTriggerEvent(device, deviceOperateLog))

	return deviceOperateLog, nil
}

func (l DeviceOperate) OnOperateResponse(ctx context.Context, gatewayDevice *model.Device, device *model.Device, iotMessage *pkgdevice.OperateMessage) error {
	operateLog, err := l.GetOperateLogByNumber(iotMessage.Id)
	if err == nil {
		if device.ID != operateLog.DeviceId {
			return exception.NewResponseError("设备不匹配")
		}
		err = l.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
			operateLog.ResponsePayload = iotMessage.Payload
			operateLog.ResponseTime = time.Unix(iotMessage.Timestamp, 0).Format(pkgmodel.TimeFormat)
			if !repository.Repository.DeviceOperateLog.UpdateDeviceOperateLog(tx, operateLog) {
				return exception.NewResponseError("更新操作记录失败")
			}

			status, exists := operateLog.ResponsePayload["cur_status"]
			if exists && reflect.TypeOf(status).Name() == "string" {
				device.DeviceCurStatus = status.(string)
				if !repository.Repository.Device.UpdateDevice(tx, device) {
					return exception.NewResponseError("更新设备失败")
				}
			}

			return nil
		})

		if err == nil {
			app.GApp.GetEvent().Publish(reflect.TypeOf(event.DeviceOperateReply{}).Name(), event.NewDeviceOperateReplyEvent(device, operateLog, iotMessage))
		}
	}

	return err
}

func (l DeviceOperate) UpdateOperateLog(operateLog *model.DeviceOperateLog) error {
	if !repository.Repository.DeviceOperateLog.UpdateDeviceOperateLog(l.GetDefaultDb(), operateLog) {
		return exception.NewResponseError("更新操作记录失败")
	}

	return nil
}

func (l DeviceOperate) GetDeviceOperateLogByNumber(device *model.Device, operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog, _ := l.GetOperateLogByNumber(operateNumber)
	if operateLog.DeviceId != device.ID {
		return nil, exception.NewResponseError("非法操作")
	}

	return operateLog, nil
}

func (l DeviceOperate) GetOperateLogByNumber(operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog := repository.Repository.DeviceOperateLog.GetDeviceOperateLogByOperateNumber(l.GetDefaultDb(), operateNumber)
	if operateLog == nil {
		return nil, exception.NewResponseError("设备操作记录不存在")
	}

	return operateLog, nil
}

func (l DeviceOperate) GetDeviceOperates(device *model.Device, page uint, pageSize uint) *pkgrepository.PageModel {
	return repository.Repository.DeviceOperateLog.GetDeviceOperates(l.GetDefaultDb(), device.ID, page, pageSize)
}

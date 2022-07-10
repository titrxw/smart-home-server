package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"reflect"
	"time"

	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
)

type DeviceOperateLogic struct {
	LogicAbstract
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
		Type:           device.Type,
		Source:         global.FApp.Name,
		OperateName:    string(operate),
		OperateNumber:  helper.Sha1(device.App.AppId + helper.UUid()),
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

func (deviceOperateLogic DeviceOperateLogic) OnOperateResponse(replyMessage model.DeviceOperateReplyMessage) (*model.Device, *model.DeviceOperateLog, error) {
	operateLog, err := deviceOperateLogic.GetOperateLogResultByNumber(replyMessage.OperateId)
	var device *model.Device = nil
	if err == nil {
		err = deviceOperateLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
			operateLog.ResponsePayload = replyMessage.PayLoad
			operateLog.ResponseTime = time.Now().Format(model.TimeFormat)
			if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(tx, operateLog) {
				return errors.New("更新操作记录失败")
			}

			status, exists := operateLog.ResponsePayload["cur_status"]
			if exists && reflect.TypeOf(status).Name() == "string" {
				device = Logic.DeviceLogic.GetDeviceById(operateLog.DeviceId)
				if device == nil {
					return errors.New("设备不存在")
				}
				device.DeviceCurStatus = status.(string)
				if !repository.Repository.DeviceRepository.UpdateDevice(tx, device) {
					return errors.New("更新设备失败")
				}
			}

			return nil
		})
	}

	return device, operateLog, err
}

func (deviceOperateLogic DeviceOperateLogic) UpdateOperateLog(operateLog *model.DeviceOperateLog) error {
	if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(deviceOperateLogic.GetDefaultDb(), operateLog) {
		return errors.New("更新操作记录失败")
	}

	return nil
}

func (deviceOperateLogic DeviceOperateLogic) GetDeviceOperateLogResultByNumber(device *model.Device, operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog, _ := deviceOperateLogic.GetOperateLogResultByNumber(operateNumber)
	if operateLog.DeviceId != device.ID {
		return nil, errors.New("非法操作")
	}

	return operateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) GetOperateLogResultByNumber(operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog := repository.Repository.DeviceOperateLogRepository.GetDeviceOperateLogByOperateNumber(deviceOperateLogic.GetDefaultDb(), operateNumber)
	if operateLog == nil {
		return nil, errors.New("设备操作记录不存在")
	}

	return operateLog, nil
}

func (deviceOperateLogic DeviceOperateLogic) GetDeviceOperates(device *model.Device, page uint, pageSize uint) *repository.PageModel {
	return repository.DeviceOperateLogRepository{}.GetDeviceOperates(deviceOperateLogic.GetDefaultDb(), device.ID, page, pageSize)
}
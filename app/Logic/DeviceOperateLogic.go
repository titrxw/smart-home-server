package logic

import (
	"context"
	"errors"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
	"reflect"
	"time"
)

type DeviceOperateLogic struct {
	LogicAbstract
}

func (this DeviceOperateLogic) IsSupportOperate(device *model.Device, operate model.DeviceOperateType) bool {
	_, exists := Logic.DeviceLogic.GetDeviceSupportMap()[device.Type]
	if !exists {
		return false
	}

	for _, element := range Logic.DeviceLogic.GetDeviceSupportMap()[device.Type].SupportOperate {
		if operate == model.DeviceOperateType(element) {
			return true
		}
	}
	return false
}

func (this DeviceOperateLogic) TriggerOperate(ctx context.Context, device *model.Device, operate model.DeviceOperateType, payload model.OperatePayload, operateLevel uint8) (*model.DeviceOperateLog, error) {
	if device.IsDelete() {
		return nil, errors.New("该设备已删除")
	}
	if device.IsDisable() {
		return nil, errors.New("该设备状态异常")
	}
	if !this.IsSupportOperate(device, operate) {
		return nil, errors.New("该设备不支持当前操作")
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

	if !repository.Repository.DeviceOperateLogRepository.AddDeviceOperateLog(this.GetDefaultDb(), deviceOperateLog) {
		return nil, errors.New("添加操作记录失败")
	}

	err = Logic.EmqxLogic.PubClientOperate(ctx, device, deviceOperateLog)
	if err != nil {
		deviceOperateLog.ResponsePayload = model.OperatePayload{"error": err.Error()}
		if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(this.GetDefaultDb(), deviceOperateLog) {
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

func (this DeviceOperateLogic) UpdateOperateLog(operateLog *model.DeviceOperateLog) error {
	if !repository.Repository.DeviceOperateLogRepository.UpdateDeviceOperateLog(this.GetDefaultDb(), operateLog) {
		return errors.New("更新操作记录失败")
	}

	return nil
}

func (this DeviceOperateLogic) GetDeviceOperateLogResultByNumber(device *model.Device, operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog := repository.Repository.DeviceOperateLogRepository.GetDeviceOperateLogByOperateNumber(this.GetDefaultDb(), operateNumber)
	if operateLog == nil {
		return nil, errors.New("设备操作记录不存在")
	}
	if operateLog.DeviceId != device.ID {
		return nil, errors.New("非法操作")
	}

	return operateLog, nil
}

func (this DeviceOperateLogic) GetOperateLogResultByNumber(operateNumber string) (*model.DeviceOperateLog, error) {
	operateLog := repository.Repository.DeviceOperateLogRepository.GetDeviceOperateLogByOperateNumber(this.GetDefaultDb(), operateNumber)
	if operateLog == nil {
		return nil, errors.New("设备操作记录不存在")
	}

	return operateLog, nil
}

func (this DeviceOperateLogic) GetDeviceOperates(device *model.Device, page uint, pageSize uint) *repository.PageModel {
	return repository.DeviceOperateLogRepository{}.GetDeviceOperates(this.GetDefaultDb(), device.ID, page, pageSize)
}

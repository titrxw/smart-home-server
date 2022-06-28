package device

import (
	"context"
	"errors"
	"github.com/titrxw/smart-home-server/app/Adapter/Interface"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"reflect"
)

type LightDeviceAdapter struct {
	Interface.DeviceAdapterInterface
}

func (this LightDeviceAdapter) GetDeviceType() string {
	return "light"
}

func (this LightDeviceAdapter) BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (this LightDeviceAdapter) AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (this LightDeviceAdapter) OnOperateResponse(deviceOperateLog *model.DeviceOperateLog) error {
	device := logic.Logic.DeviceLogic.GetDeviceById(deviceOperateLog.DeviceId)
	if device == nil {
		return errors.New("设备不存在")
	}

	status, exists := deviceOperateLog.ResponsePayload["cur_status"]
	if exists && reflect.TypeOf(status).Name() == "string" {
		device.DeviceCurStatus = status.(string)
		return logic.Logic.DeviceLogic.UpdateUserDevice(context.Background(), device)
	}
	return nil
}

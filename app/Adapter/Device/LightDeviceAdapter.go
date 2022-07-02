package device

import (
	"context"
	"errors"
	"reflect"

	"github.com/titrxw/smart-home-server/app/Adapter/Interface"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type LightDeviceAdapter struct {
	Interface.DeviceAdapterInterface
}

func (lightAdapter LightDeviceAdapter) GetDeviceType() string {
	return "light"
}

func (lightAdapter LightDeviceAdapter) BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) OnOperateResponse(deviceOperateLog *model.DeviceOperateLog) error {
	device := logic.Logic.DeviceLogic.GetDeviceById(deviceOperateLog.DeviceId)
	if device == nil {
		return errors.New("设备不存在")
	}

	status, exists := deviceOperateLog.ResponsePayload["cur_status"]
	if exists && reflect.TypeOf(status).Name() == "string" {
		device.DeviceCurStatus = status.(string)
		return logic.Logic.DeviceLogic.UpdateDevice(context.Background(), device)
	}
	return nil
}

func (lightAdapter LightDeviceAdapter) OnReport(device *model.Device, payload string) error {
	return nil
}

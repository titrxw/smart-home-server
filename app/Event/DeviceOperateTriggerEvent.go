package event

import model "github.com/titrxw/smart-home-server/app/Model"

type DeviceOperateTriggerEvent struct {
	Device           *model.Device
	DeviceOperateLog *model.DeviceOperateLog
}

func NewDeviceOperateTriggerEvent(device *model.Device, deviceOperateLog *model.DeviceOperateLog) DeviceOperateTriggerEvent {
	return DeviceOperateTriggerEvent{
		Device:           device,
		DeviceOperateLog: deviceOperateLog,
	}
}

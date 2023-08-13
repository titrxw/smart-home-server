package event

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
)

type DeviceOperateTrigger struct {
	Device           *model.Device
	DeviceOperateLog *model.DeviceOperateLog
}

func NewDeviceOperateTriggerEvent(device *model.Device, deviceOperateLog *model.DeviceOperateLog) DeviceOperateTrigger {
	return DeviceOperateTrigger{
		Device:           device,
		DeviceOperateLog: deviceOperateLog,
	}
}

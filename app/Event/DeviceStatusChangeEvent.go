package event

import (
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceStatusChangeEvent struct {
	Device *model.Device
}

func NewDeviceStatusChangeEvent(device *model.Device) DeviceStatusChangeEvent {
	return DeviceStatusChangeEvent{
		Device: device,
	}
}

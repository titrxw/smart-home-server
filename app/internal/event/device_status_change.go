package event

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
)

type DeviceStatusChange struct {
	Device *model.Device
}

func NewDeviceStatusChangeEvent(device *model.Device) DeviceStatusChange {
	return DeviceStatusChange{
		Device: device,
	}
}

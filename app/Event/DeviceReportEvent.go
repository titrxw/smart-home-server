package event

import model "github.com/titrxw/smart-home-server/app/Model"

type DeviceReportEvent struct {
	Device  *model.Device
	Payload string
}

func NewDeviceReportEvent(device *model.Device, payload string) DeviceReportEvent {
	return DeviceReportEvent{
		Device:  device,
		Payload: payload,
	}
}

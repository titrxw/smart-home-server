package event

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceReportEvent struct {
	Device     *model.Device
	CloudEvent *cloudevents.Event
}

func NewDeviceReportEvent(device *model.Device, cloudEvent *cloudevents.Event) DeviceReportEvent {
	return DeviceReportEvent{
		Device:     device,
		CloudEvent: cloudEvent,
	}
}

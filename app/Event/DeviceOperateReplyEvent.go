package event

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceOperateReplyEvent struct {
	Device           *model.Device
	DeviceOperateLog *model.DeviceOperateLog
	CloudEvent       *cloudevents.Event
}

func NewDeviceOperateReplyEvent(device *model.Device, deviceOperateLog *model.DeviceOperateLog, cloudEvent *cloudevents.Event) DeviceOperateReplyEvent {
	return DeviceOperateReplyEvent{
		Device:           device,
		DeviceOperateLog: deviceOperateLog,
		CloudEvent:       cloudEvent,
	}
}

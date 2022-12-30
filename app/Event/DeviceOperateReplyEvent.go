package event

import (
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceOperateReplyEvent struct {
	Device           *model.Device
	DeviceOperateLog *model.DeviceOperateLog
	Message          *model.IotMessage
}

func NewDeviceOperateReplyEvent(device *model.Device, deviceOperateLog *model.DeviceOperateLog, message *model.IotMessage) DeviceOperateReplyEvent {
	return DeviceOperateReplyEvent{
		Device:           device,
		DeviceOperateLog: deviceOperateLog,
		Message:          message,
	}
}

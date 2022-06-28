package event

import (
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceOperateReplyEvent struct {
	DeviceOperateLog *model.DeviceOperateLog
}

func NewDeviceOperateReplyEvent(deviceOperateLog *model.DeviceOperateLog) DeviceOperateReplyEvent {
	return DeviceOperateReplyEvent{
		DeviceOperateLog: deviceOperateLog,
	}
}

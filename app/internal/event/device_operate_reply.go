package event

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/pkg/device"
)

type DeviceOperateReply struct {
	Device           *model.Device
	DeviceOperateLog *model.DeviceOperateLog
	Message          *device.OperateMessage
}

func NewDeviceOperateReplyEvent(device *model.Device, deviceOperateLog *model.DeviceOperateLog, message *device.OperateMessage) DeviceOperateReply {
	return DeviceOperateReply{
		Device:           device,
		DeviceOperateLog: deviceOperateLog,
		Message:          message,
	}
}

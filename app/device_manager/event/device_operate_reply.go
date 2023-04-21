package event

import (
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	_interface "github.com/titrxw/smart-home-server/app/devices/interface"
)

type DeviceOperateReply struct {
	Device           *model.Device
	DeviceOperateLog *model.DeviceOperateLog
	Message          *_interface.DeviceOperateMessage
}

func NewDeviceOperateReplyEvent(device *model.Device, deviceOperateLog *model.DeviceOperateLog, message *_interface.DeviceOperateMessage) DeviceOperateReply {
	return DeviceOperateReply{
		Device:           device,
		DeviceOperateLog: deviceOperateLog,
		Message:          message,
	}
}

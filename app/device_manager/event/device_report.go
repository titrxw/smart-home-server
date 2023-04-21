package event

import (
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	_interface "github.com/titrxw/smart-home-server/app/devices/interface"
)

type DeviceReport struct {
	Device    *model.Device
	ReportLog *model.DeviceReportLog
	Message   *_interface.DeviceOperateMessage
}

func NewDeviceReportEvent(device *model.Device, reportLog *model.DeviceReportLog, message *_interface.DeviceOperateMessage) DeviceReport {
	return DeviceReport{
		Device:    device,
		ReportLog: reportLog,
		Message:   message,
	}
}

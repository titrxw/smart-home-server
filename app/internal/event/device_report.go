package event

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/pkg/device"
)

type DeviceReport struct {
	Device    *model.Device
	ReportLog *model.DeviceReportLog
	Message   *device.OperateMessage
}

func NewDeviceReportEvent(device *model.Device, reportLog *model.DeviceReportLog, message *device.OperateMessage) DeviceReport {
	return DeviceReport{
		Device:    device,
		ReportLog: reportLog,
		Message:   message,
	}
}

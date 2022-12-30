package event

import (
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceReportEvent struct {
	Device    *model.Device
	ReportLog *model.DeviceReportLog
	Message   *model.IotMessage
}

func NewDeviceReportEvent(device *model.Device, reportLog *model.DeviceReportLog, message *model.IotMessage) DeviceReportEvent {
	return DeviceReportEvent{
		Device:    device,
		ReportLog: reportLog,
		Message:   message,
	}
}

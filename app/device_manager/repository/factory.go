package repository

type Factory struct {
	App
	AppProxy
	Device
	DeviceOperateLog
	DeviceReportLog
	Setting
	User
}

var Repository = new(Factory)

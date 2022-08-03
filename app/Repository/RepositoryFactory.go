package repository

type RepositoryFactory struct {
	AppRepository
	DeviceRepository
	DeviceOperateLogRepository
	DeviceReportLogRepository
	SettingRepository
	UserRepository
}

var Repository = new(RepositoryFactory)

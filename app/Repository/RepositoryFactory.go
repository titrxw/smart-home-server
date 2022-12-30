package repository

type RepositoryFactory struct {
	AppRepository
	AppProxyRepository
	DeviceRepository
	DeviceOperateLogRepository
	DeviceReportLogRepository
	SettingRepository
	UserRepository
}

var Repository = new(RepositoryFactory)

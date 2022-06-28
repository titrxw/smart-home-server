package repository

type RepositoryFactory struct {
	AppRepository
	DeviceRepository
	DeviceOperateLogRepository
	SettingRepository
	UserRepository
}

var Repository = new(RepositoryFactory)

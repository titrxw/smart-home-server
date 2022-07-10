package repository

import (
	"time"

	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	RepositoryAbstract
}

func (deviceRepository DeviceRepository) AddUserDeviceByApp(db *gorm.DB, userId model.UID, app *model.App, device *model.Device) *model.Device {
	device.UserId = userId
	device.AppId = app.ID
	device.CreatedAt = model.LocalTime(time.Now())

	result := db.Create(device)
	if result.RowsAffected != 1 {
		return nil
	}

	device.App = app
	return device
}

func (deviceRepository DeviceRepository) GetUserDevices(db *gorm.DB, userId model.UID, page uint, pageSize uint) *PageModel {
	devices := make([]model.Device, 0)
	pageData := &PageModel{
		CurPage:  page,
		Total:    0,
		PageSize: pageSize,
		Data:     &devices,
	}

	totalQuery := db.Model(&model.Device{}).Where("user_id = ?", userId).Where("device_status != ?", model.DEVICE_DELETE)
	var total int64
	totalQuery.Count(&total)
	if total > 0 {
		totalQuery.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&devices)
		pageData.Total = uint64(total)
	}

	return pageData
}

func (deviceRepository DeviceRepository) GetUserDeviceById(db *gorm.DB, userId model.UID, id uint) *model.Device {
	device := new(model.Device)
	result := db.Where("id = ?", id).Where("user_id = ?", userId).Where("device_status != ?", model.DEVICE_DELETE).First(device)
	if result.RowsAffected == 1 {
		device.App = Repository.AppRepository.GetById(db, device.AppId)
		if device.App == nil {
			device = nil
		}
		return device
	}

	return nil
}

func (deviceRepository DeviceRepository) GetDeviceByApp(db *gorm.DB, app *model.App) *model.Device {
	device := new(model.Device)
	result := db.Where("app_id = ?", app.ID).Where("device_status != ?", model.DEVICE_DELETE).First(device)
	if result.RowsAffected == 1 {
		device.App = app
		return device
	}

	return nil
}

func (deviceRepository DeviceRepository) GetDeviceById(db *gorm.DB, id uint) *model.Device {
	device := new(model.Device)
	result := db.Where("id = ?", id).Where("device_status != ?", model.DEVICE_DELETE).First(device)
	if result.RowsAffected == 1 {
		device.App = Repository.AppRepository.GetById(db, device.AppId)
		if device.App == nil {
			device = nil
		}
		return device
	}

	return nil
}

func (deviceRepository DeviceRepository) UpdateDevice(db *gorm.DB, device *model.Device) bool {
	result := db.Save(device)
	return result.Error == nil
}

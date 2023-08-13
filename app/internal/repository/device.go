package repository

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
	pkgmodel "github.com/titrxw/smart-home-server/app/pkg/model"
	"github.com/titrxw/smart-home-server/app/pkg/repository"
	"time"

	"gorm.io/gorm"
)

type Device struct {
	repository.Abstract
}

func (r Device) AddUserDeviceByApp(db *gorm.DB, userId model.UID, app *model.App, device *model.Device) *model.Device {
	device.UserId = userId
	device.AppId = app.ID
	device.DeviceAppId = app.AppId
	device.CreatedAt = pkgmodel.LocalTime(time.Now())

	result := db.Create(device)
	if result.RowsAffected != 1 {
		return nil
	}

	device.App = app

	return device
}

func (r Device) GetUserDevices(db *gorm.DB, userId model.UID, page uint, pageSize uint) *repository.PageModel {
	devices := make([]model.Device, 0)
	pageData := &repository.PageModel{
		CurPage:  page,
		Total:    0,
		PageSize: pageSize,
		Data:     &devices,
	}

	totalQuery := db.Model(&model.Device{}).Where("user_id = ?", userId).Where("device_status != ?", model.DeviceDelete)
	var total int64
	totalQuery.Count(&total)
	if total > 0 {
		totalQuery.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&devices)
		pageData.Total = uint64(total)
	}

	return pageData
}

func (r Device) GetUserDeviceById(db *gorm.DB, userId model.UID, id uint) *model.Device {
	device := new(model.Device)
	result := db.Where("id = ?", id).Where("user_id = ?", userId).Where("device_status != ?", model.DeviceDelete).First(device)
	if result.RowsAffected == 1 {
		device.App = Repository.App.GetById(db, device.AppId)
		if device.App == nil {
			device = nil
		}
		return device
	}

	return nil
}

func (r Device) GetUserDeviceByApp(db *gorm.DB, userId model.UID, app *model.App) *model.Device {
	device := new(model.Device)
	result := db.Where("user_id = ?", userId).Where("app_id = ?", app.ID).Where("device_status != ?", model.DeviceDelete).First(device)
	if result.RowsAffected == 1 {
		device.App = Repository.App.GetById(db, device.AppId)
		if device.App == nil {
			device = nil
		}
		return device
	}

	return nil
}

func (r Device) GetDeviceByApp(db *gorm.DB, app *model.App) *model.Device {
	device := new(model.Device)
	result := db.Where("app_id = ?", app.ID).Where("device_status != ?", model.DeviceDelete).First(device)
	if result.RowsAffected == 1 {
		device.App = app
		return device
	}

	return nil
}

func (r Device) GetDeviceById(db *gorm.DB, id uint) *model.Device {
	device := new(model.Device)
	result := db.Where("id = ?", id).Where("device_status != ?", model.DeviceDelete).First(device)
	if result.RowsAffected == 1 {
		device.App = Repository.App.GetById(db, device.AppId)
		if device.App == nil {
			device = nil
		}
		return device
	}

	return nil
}

func (r Device) UpdateDevice(db *gorm.DB, device *model.Device) bool {
	result := db.Save(device)

	return result.Error == nil
}

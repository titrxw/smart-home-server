package repository

import (
	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
)

type DeviceOperateLogRepository struct {
	RepositoryAbstract
}

func (deviceOperateLogRepository DeviceOperateLogRepository) AddDeviceOperateLog(db *gorm.DB, operateLog *model.DeviceOperateLog) bool {
	result := db.Create(operateLog)

	return result.Error == nil
}

func (deviceOperateLogRepository DeviceOperateLogRepository) UpdateDeviceOperateLog(db *gorm.DB, operateLog *model.DeviceOperateLog) bool {
	result := db.Save(operateLog)

	return result.Error == nil
}

func (deviceOperateLogRepository DeviceOperateLogRepository) GetDeviceOperateLogByOperateNumber(db *gorm.DB, operateNumber string) *model.DeviceOperateLog {
	deviceOperateLog := new(model.DeviceOperateLog)
	result := db.Where("operate_number = ?", operateNumber).First(deviceOperateLog)
	if result.RowsAffected == 1 {
		return deviceOperateLog
	}

	return nil
}

func (deviceOperateLogRepository DeviceOperateLogRepository) GetDeviceOperates(db *gorm.DB, deviceId uint, page uint, pageSize uint) *PageModel {
	deviceOperates := make([]model.DeviceOperateLog, 0)
	pageData := &PageModel{
		CurPage:  page,
		Total:    0,
		PageSize: pageSize,
		Data:     &deviceOperates,
	}

	totalQuery := db.Model(&model.DeviceOperateLog{}).Where("device_id = ?", deviceId)
	var total int64
	totalQuery.Count(&total)
	if total > 0 {
		totalQuery.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&deviceOperates)
		pageData.Total = uint64(total)
	}

	return pageData
}

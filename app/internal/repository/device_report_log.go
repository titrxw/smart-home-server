package repository

import (
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/pkg/repository"
	"gorm.io/gorm"
)

type DeviceReportLog struct {
	repository.Abstract
}

func (r DeviceReportLog) AddDeviceReportLog(db *gorm.DB, reportLog *model.DeviceReportLog) bool {
	result := db.Create(reportLog)

	return result.Error == nil
}

func (r DeviceReportLog) GetDeviceReportLogByReportNumber(db *gorm.DB, reportNumber string) *model.DeviceReportLog {
	deviceReportLog := new(model.DeviceReportLog)
	result := db.Where("report_number = ?", reportNumber).First(deviceReportLog)
	if result.RowsAffected == 1 {
		return deviceReportLog
	}

	return nil
}

func (r DeviceReportLog) GetDeviceReports(db *gorm.DB, deviceId uint, page uint, pageSize uint) *repository.PageModel {
	deviceReports := make([]model.DeviceReportLog, 0)
	pageData := &repository.PageModel{
		CurPage:  page,
		Total:    0,
		PageSize: pageSize,
		Data:     &deviceReports,
	}

	totalQuery := db.Model(&model.DeviceReportLog{}).Where("device_id = ?", deviceId)
	var total int64
	totalQuery.Count(&total)
	if total > 0 {
		totalQuery.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&deviceReports)
		pageData.Total = uint64(total)
	}

	return pageData
}

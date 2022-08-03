package repository

import (
	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
)

type DeviceReportLogRepository struct {
	RepositoryAbstract
}

func (deviceReportLogRepository DeviceReportLogRepository) AddDeviceReportLog(db *gorm.DB, reportLog *model.DeviceReportLog) bool {
	result := db.Create(reportLog)

	return result.Error == nil
}

func (deviceReportLogRepository DeviceReportLogRepository) GetDeviceReportLogByReportNumber(db *gorm.DB, reportNumber string) *model.DeviceReportLog {
	deviceReportLog := new(model.DeviceReportLog)
	result := db.Where("report_number = ?", reportNumber).First(deviceReportLog)
	if result.RowsAffected == 1 {
		return deviceReportLog
	}

	return nil
}

func (deviceReportLogRepository DeviceReportLogRepository) GetDeviceReports(db *gorm.DB, deviceId uint, page uint, pageSize uint) *PageModel {
	deviceReports := make([]model.DeviceReportLog, 0)
	pageData := &PageModel{
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

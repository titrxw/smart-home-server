package logic

import (
	"encoding/json"
	"errors"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	global "github.com/titrxw/go-framework/src/Global"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
	"time"
)

type DeviceReportLogic struct {
	LogicAbstract
}

func (deviceReportLogic DeviceReportLogic) OnReport(device *model.Device, cloudEvent *cloudevents.Event) (*model.DeviceReportLog, error) {
	if !Logic.DeviceLogic.IsSupportReport(device, model.DeviceReportType(cloudEvent.Type())) {
		return nil, errors.New("report 不支持")
	}

	payLoad := model.ReportPayload{}
	err := json.Unmarshal(cloudEvent.Data(), &payLoad)
	if err != nil {
		return nil, err
	}

	payLoad["report_id"] = cloudEvent.ID()
	deviceReportLog := &model.DeviceReportLog{
		DeviceId:      device.ID,
		DeviceType:    device.Type,
		Source:        global.FApp.Name,
		ReportName:    cloudEvent.Type(),
		ReportNumber:  helper.Sha1(device.App.AppId + helper.UUid()),
		ReportTime:    model.LocalTime(time.Now()),
		ReportPayload: payLoad,
		ReportLevel:   2,
		CreatedAt:     model.LocalTime(time.Now()),
	}
	if !repository.Repository.DeviceReportLogRepository.AddDeviceReportLog(deviceReportLogic.GetDefaultDb(), deviceReportLog) {
		return nil, errors.New("添加上报记录失败")
	}

	return deviceReportLog, nil
}

func (deviceReportLogic DeviceReportLogic) GetDeviceReportLogByNumber(device *model.Device, reportNumber string) (*model.DeviceReportLog, error) {
	reportLog := repository.Repository.DeviceReportLogRepository.GetDeviceReportLogByReportNumber(deviceReportLogic.GetDefaultDb(), reportNumber)
	if reportLog == nil {
		return nil, errors.New("设备上报记录不存在")
	}

	return reportLog, nil
}

func (deviceReportLogic DeviceReportLogic) GetDeviceReports(device *model.Device, page uint, pageSize uint) *repository.PageModel {
	return repository.Repository.DeviceReportLogRepository.GetDeviceReports(deviceReportLogic.GetDefaultDb(), device.ID, page, pageSize)
}

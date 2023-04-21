package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/event"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/titrxw/smart-home-server/app/device_manager/repository"
	_interface "github.com/titrxw/smart-home-server/app/devices/interface"
	"github.com/titrxw/smart-home-server/app/devices/manager"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"reflect"
	"time"
)

type DeviceReport struct {
	Abstract
}

func (l DeviceReport) OnReport(gatewayDevice *model.Device, device *model.Device, iotMessage *_interface.DeviceOperateMessage) error {
	if !Logic.Device.IsSupportReport(device, model.DeviceReportType(iotMessage.EventType)) {
		return exception.NewResponseError("report 不支持")
	}
	if iotMessage.Id == device.App.AppId {
		iotMessage.Id = Logic.DeviceReport.GetOperateOrReportNumber(device.App.AppId)
	}

	deviceReportLog := &model.DeviceReportLog{
		DeviceId:        device.ID,
		DeviceGatewayId: gatewayDevice.ID,
		DeviceType:      device.TypeName,
		Source:          helper.GetAppName(),
		ReportName:      iotMessage.EventType,
		ReportNumber:    helper.Sha1(device.App.AppId + helper.UUid()),
		ReportTime:      model.LocalTime(time.Now()),
		ReportPayload:   model.ReportPayload(iotMessage.Payload),
		ReportLevel:     2,
		CreatedAt:       model.LocalTime(time.Unix(iotMessage.Timestamp, 0)),
	}
	if !repository.Repository.DeviceReportLog.AddDeviceReportLog(l.GetDefaultDb(), deviceReportLog) {
		return exception.NewResponseError("添加上报记录失败")
	}

	err := manager.GetDevice(gatewayDevice.TypeName).OnReport(context.Background(), gatewayDevice.App.AppId, device.App.AppId, device.TypeName, iotMessage)
	app.GApp.GetEvent().Publish(reflect.TypeOf(event.DeviceReport{}).Name(), event.NewDeviceReportEvent(device, deviceReportLog, iotMessage))

	return err
}

func (l DeviceReport) GetDeviceReportLogByNumber(device *model.Device, reportNumber string) (*model.DeviceReportLog, error) {
	reportLog := repository.Repository.DeviceReportLog.GetDeviceReportLogByReportNumber(l.GetDefaultDb(), reportNumber)
	if reportLog == nil {
		return nil, exception.NewResponseError("设备上报记录不存在")
	}
	if reportLog.DeviceId != device.ID {
		return nil, exception.NewResponseError("非法操作")
	}

	return reportLog, nil
}

func (l DeviceReport) GetDeviceReports(device *model.Device, page uint, pageSize uint) *repository.PageModel {
	return repository.Repository.DeviceReportLog.GetDeviceReports(l.GetDefaultDb(), device.ID, page, pageSize)
}

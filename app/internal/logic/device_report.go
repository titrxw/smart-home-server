package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/internal/event"
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/internal/repository"
	"github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
	pkgmodel "github.com/titrxw/smart-home-server/app/pkg/model"
	pkgrepository "github.com/titrxw/smart-home-server/app/pkg/repository"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"reflect"
	"time"
)

type DeviceReport struct {
	logic.Abstract
}

func (l DeviceReport) OnReport(ctx context.Context, gatewayDevice *model.Device, device *model.Device, iotMessage *device.OperateMessage) error {
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
		ReportTime:      pkgmodel.LocalTime(time.Now()),
		ReportPayload:   model.ReportPayload(iotMessage.Payload),
		ReportLevel:     2,
		CreatedAt:       pkgmodel.LocalTime(time.Unix(iotMessage.Timestamp, 0)),
	}
	if !repository.Repository.DeviceReportLog.AddDeviceReportLog(l.GetDefaultDb(), deviceReportLog) {
		return exception.NewResponseError("添加上报记录失败")
	}

	app.GApp.GetEvent().Publish(reflect.TypeOf(event.DeviceReport{}).Name(), event.NewDeviceReportEvent(device, deviceReportLog, iotMessage))

	return nil
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

func (l DeviceReport) GetDeviceReports(device *model.Device, page uint, pageSize uint) *pkgrepository.PageModel {
	return repository.Repository.DeviceReportLog.GetDeviceReports(l.GetDefaultDb(), device.ID, page, pageSize)
}

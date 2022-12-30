package logic

import (
	"context"
	"reflect"
	"time"

	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
)

type DeviceReportLogic struct {
	LogicAbstract
}

func (deviceReportLogic DeviceReportLogic) OnReport(gatewayDevice *model.Device, device *model.Device, iotMessage *model.IotMessage) error {
	if !Logic.DeviceLogic.IsSupportReport(device, model.DeviceReportType(iotMessage.EventType)) {
		return exception.NewLogicError("report 不支持")
	}
	if iotMessage.Id == device.App.AppId {
		iotMessage.Id = Logic.DeviceReportLogic.GetOperateOrReportNumber(device.App.AppId)
	}

	deviceReportLog := &model.DeviceReportLog{
		DeviceId:        device.ID,
		DeviceGatewayId: gatewayDevice.ID,
		DeviceType:      device.TypeName,
		Source:          global.FApp.Name,
		ReportName:      iotMessage.EventType,
		ReportNumber:    helper.Sha1(device.App.AppId + helper.UUid()),
		ReportTime:      model.LocalTime(time.Now()),
		ReportPayload:   model.ReportPayload(iotMessage.Payload),
		ReportLevel:     2,
		CreatedAt:       model.LocalTime(time.Unix(iotMessage.Timestamp, 0)),
	}
	if !repository.Repository.DeviceReportLogRepository.AddDeviceReportLog(deviceReportLogic.GetDefaultDb(), deviceReportLog) {
		return exception.NewLogicError("添加上报记录失败")
	}

	err := Logic.DeviceLogic.GetDeviceAdapter(gatewayDevice.TypeName).OnReport(context.Background(), gatewayDevice, device, deviceReportLog, iotMessage)
	global.FApp.Event.Publish(reflect.TypeOf(event.DeviceReportEvent{}).Name(), event.NewDeviceReportEvent(device, deviceReportLog, iotMessage))

	return err
}

func (deviceReportLogic DeviceReportLogic) GetDeviceReportLogByNumber(device *model.Device, reportNumber string) (*model.DeviceReportLog, error) {
	reportLog := repository.Repository.DeviceReportLogRepository.GetDeviceReportLogByReportNumber(deviceReportLogic.GetDefaultDb(), reportNumber)
	if reportLog == nil {
		return nil, exception.NewLogicError("设备上报记录不存在")
	}
	if reportLog.DeviceId != device.ID {
		return nil, exception.NewLogicError("非法操作")
	}

	return reportLog, nil
}

func (deviceReportLogic DeviceReportLogic) GetDeviceReports(device *model.Device, page uint, pageSize uint) *repository.PageModel {
	return repository.Repository.DeviceReportLogRepository.GetDeviceReports(deviceReportLogic.GetDefaultDb(), device.ID, page, pageSize)
}

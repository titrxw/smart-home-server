package Interface

import (
	"context"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
)

type DeviceAdapterInterface interface {
	GetDeviceConfig() config.Device
	GetCtrlTopic(appId string, componentAppId string) string
	GetReportTopic(appId string) string
	GetAvailabilityTopic(appId string) string
	BeforeTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	AfterTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	OnOperateResponse(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog, message *model.IotMessage) error
	OnReport(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceReportLog *model.DeviceReportLog, message *model.IotMessage) error
	PackMessage(message *model.IotMessage) (string, error)
	UnPackMessage(payload string) (*model.IotMessage, error)
}

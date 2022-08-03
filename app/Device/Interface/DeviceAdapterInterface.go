package Interface

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
)

type DeviceAdapterInterface interface {
	GetDeviceConfig() config.Device
	BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	OnOperateResponse(device *model.Device, deviceOperateLog *model.DeviceOperateLog, cloudEvent *cloudevents.Event) error
	OnReport(device *model.Device, deviceReportLog *model.DeviceReportLog, cloudEvent *cloudevents.Event) error
}

package Interface

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
)

type DeviceAdapterInterface interface {
	GetDeviceConfig() config.Device
	BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	OnOperateResponse(client mqtt.Client, device *model.Device, deviceOperateLog *model.DeviceOperateLog, cloudEvent *cloudevents.Event) error
	OnReport(client mqtt.Client, device *model.Device, deviceReportLog *model.DeviceReportLog, cloudEvent *cloudevents.Event) error
}

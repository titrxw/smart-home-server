package light

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
)

type LightDeviceAdapter struct {
	Interface.DeviceAdapterInterface
}

func (lightAdapter LightDeviceAdapter) GetDeviceConfig() config.Device {
	return config.Device{
		Type:           "light",
		Name:           "电灯",
		SupportOperate: []string{"on", "off"},
		OperateDesc:    map[string]string{"on": "开灯", "off": "关灯"},
		SupportReport:  nil,
		Setting:        nil,
	}
}

func (lightAdapter LightDeviceAdapter) BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) OnOperateResponse(client mqtt.Client, device *model.Device, deviceOperateLog *model.DeviceOperateLog, cloudEvent *cloudevents.Event) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) OnReport(client mqtt.Client, device *model.Device, deviceReportLog *model.DeviceReportLog, cloudEvent *cloudevents.Event) error {
	return nil
}

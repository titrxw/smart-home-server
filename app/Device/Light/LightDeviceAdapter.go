package light

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
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
	}
}

func (lightAdapter LightDeviceAdapter) BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) OnOperateResponse(device *model.Device, deviceOperateLog *model.DeviceOperateLog, cloudEvent *cloudevents.Event) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) OnReport(device *model.Device, cloudEvent *cloudevents.Event) error {
	return nil
}

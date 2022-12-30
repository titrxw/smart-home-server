package light

import (
	"context"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
)

type LightDeviceAdapter struct {
	Interface.DeviceAdapterAbstract
}

func (lightAdapter LightDeviceAdapter) GetDeviceConfig() config.Device {
	return config.Device{
		Type:           model.DEVICE_APP_TYPE,
		TypeName:       "light",
		Name:           "电灯",
		NeedGateway:    true,
		SupportOperate: []string{"on", "off"},
		OperateDesc:    map[string]string{"on": "开灯", "off": "关灯"},
		SupportReport:  nil,
		Setting:        nil,
	}
}

func (lightAdapter LightDeviceAdapter) BeforeTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) AfterTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) OnOperateResponse(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog, message *model.IotMessage) error {
	return nil
}

func (lightAdapter LightDeviceAdapter) OnReport(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceReportLog *model.DeviceReportLog, message *model.IotMessage) error {
	return nil
}

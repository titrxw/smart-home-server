package light

import (
	"context"
	deviceInterface "github.com/titrxw/smart-home-server/app/devices/interface"
)

type DeviceAdapter struct {
	deviceInterface.Abstract
}

func (a DeviceAdapter) GetDeviceConfig() deviceInterface.Device {
	return deviceInterface.Device{
		Type:           deviceInterface.DeviceAppType,
		TypeName:       "light",
		Name:           "电灯",
		NeedGateway:    true,
		SupportOperate: []string{"on", "off"},
		OperateDesc:    map[string]string{"on": "开灯", "off": "关灯"},
		SupportReport:  nil,
		Setting:        nil,
	}
}

func (a DeviceAdapter) BeforeTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	return nil
}

func (a DeviceAdapter) AfterTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	return nil
}

func (a DeviceAdapter) OnOperateResponse(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, operatePayLoad map[string]interface{}, message *deviceInterface.DeviceOperateMessage) error {
	return nil
}

func (a DeviceAdapter) OnReport(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	return nil
}

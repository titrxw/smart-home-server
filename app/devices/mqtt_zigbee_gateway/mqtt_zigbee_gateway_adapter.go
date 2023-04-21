package mqtt_zigbee_gateway

import (
	"context"
	"github.com/titrxw/smart-home-server/app/common/helper"
	deviceInterface "github.com/titrxw/smart-home-server/app/devices/interface"
	"github.com/titrxw/smart-home-server/app/devices/manager"
)

type MqttZigbeeGatewayAdapter struct {
	deviceInterface.Abstract
}

func (a MqttZigbeeGatewayAdapter) GetDeviceConfig() deviceInterface.Device {
	return deviceInterface.Device{
		Type:           deviceInterface.DeviceGatewayAppType,
		TypeName:       "mqtt_zigbee_gateway",
		Name:           "mqtt_zigbee网关",
		NeedGateway:    false,
		SupportOperate: nil,
		OperateDesc:    nil,
		SupportReport:  nil,
		Setting:        nil,
	}
}

func (a MqttZigbeeGatewayAdapter) GetCtrlTopic(appId string, componentAppId string) string {
	return "/zigbee2mqtt/" + helper.GetAppName() + "/" + appId + "/device/" + componentAppId + "/set"
}

func (a MqttZigbeeGatewayAdapter) GetReportTopic(appId string) string {
	return "/zigbee2mqtt/" + helper.GetAppName() + "/" + appId + "/device/+/get"
}

func (a MqttZigbeeGatewayAdapter) GetAvailabilityTopic(appId string) string {
	return "/zigbee2mqtt/" + helper.GetAppName() + "/" + appId + "/device/+/availability"
}

func (a MqttZigbeeGatewayAdapter) BeforeTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	return manager.GetDevice(deviceType).BeforeTriggerOperate(ctx, gatewayDeviceAppId, deviceAppId, deviceType, message)
}

func (a MqttZigbeeGatewayAdapter) AfterTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	return manager.GetDevice(deviceType).AfterTriggerOperate(ctx, gatewayDeviceAppId, deviceAppId, deviceType, message)
}

func (a MqttZigbeeGatewayAdapter) OnOperateResponse(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, operatePayLoad map[string]interface{}, message *deviceInterface.DeviceOperateMessage) error {
	return manager.GetDevice(deviceType).OnOperateResponse(ctx, gatewayDeviceAppId, deviceAppId, deviceType, operatePayLoad, message)
}

func (a MqttZigbeeGatewayAdapter) OnReport(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	return manager.GetDevice(deviceType).OnReport(ctx, gatewayDeviceAppId, deviceAppId, deviceType, message)
}

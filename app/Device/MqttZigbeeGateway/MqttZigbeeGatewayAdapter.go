package mqtt_zigbee

import (
	"context"
	global "github.com/titrxw/go-framework/src/Global"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
)

type MqttZigbeeGatewayAdapter struct {
	Interface.DeviceAdapterAbstract
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) GetDeviceConfig() config.Device {
	return config.Device{
		Type:           model.DEVICE_GATEWAY_APP_TYPE,
		TypeName:       "mqtt_zigbee_gateway",
		Name:           "mqtt_zigbee网关",
		NeedGateway:    false,
		SupportOperate: nil,
		OperateDesc:    nil,
		SupportReport:  nil,
		Setting:        nil,
	}
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) GetCtrlTopic(appId string, componentAppId string) string {
	return "/zigbee2mqtt/" + global.FApp.Name + "/" + appId + "/device/" + componentAppId + "/set"
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) GetReportTopic(appId string) string {
	return "/zigbee2mqtt/" + global.FApp.Name + "/" + appId + "/device/+/get"
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) GetAvailabilityTopic(appId string) string {
	return "/zigbee2mqtt/" + global.FApp.Name + "/" + appId + "/device/+/availability"
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) BeforeTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return logic.Logic.DeviceLogic.GetDeviceAdapter(device.TypeName).BeforeTriggerOperate(ctx, device, gatewayDevice, deviceOperateLog)
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) AfterTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return logic.Logic.DeviceLogic.GetDeviceAdapter(device.TypeName).AfterTriggerOperate(ctx, device, gatewayDevice, deviceOperateLog)
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) OnOperateResponse(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog, message *model.IotMessage) error {
	return logic.Logic.DeviceLogic.GetDeviceAdapter(device.TypeName).OnOperateResponse(ctx, gatewayDevice, device, deviceOperateLog, message)
}

func (mqttZigbeeGatewayAdapter MqttZigbeeGatewayAdapter) OnReport(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceReportLog *model.DeviceReportLog, message *model.IotMessage) error {
	return logic.Logic.DeviceLogic.GetDeviceAdapter(device.TypeName).OnReport(ctx, gatewayDevice, device, deviceReportLog, message)
}

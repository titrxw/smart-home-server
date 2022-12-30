package mqtt_zigbee

import (
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	mqtt "github.com/titrxw/smart-home-server/app/Mqtt"
	subscribe "github.com/titrxw/smart-home-server/app/Mqtt/Subscribe"
)

type MqttZigbeeGatewayProvider struct {
	provider.ProviderAbstract
}

func (mqttZigbeeGatewayProvider *MqttZigbeeGatewayProvider) Register(options interface{}) {
	adapter := new(MqttZigbeeGatewayAdapter)
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(adapter)

	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReportSubscribe(adapter.GetReportTopic("+")))
	mqtt.GetSubscribeManager().RegisterSubscribe(NewDeviceAvailabilitySubscribe(adapter.GetAvailabilityTopic("+")))
}

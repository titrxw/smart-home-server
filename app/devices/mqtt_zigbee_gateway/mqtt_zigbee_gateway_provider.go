package mqtt_zigbee_gateway

import (
	"github.com/titrxw/smart-home-server/app/device_manager/mqtt/subscribe"
	"github.com/titrxw/smart-home-server/app/devices/manager"
	"github.com/titrxw/smart-home-server/app/mqtt"
	"github.com/we7coreteam/w7-rangine-go-support/src/provider"
)

type Provider struct {
	provider.Abstract
}

func (p *Provider) Register() {
	adapter := new(MqttZigbeeGatewayAdapter)
	manager.RegisterDevice(adapter)

	mqtt.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReportSubscribe(adapter.GetReportTopic("+")))
	mqtt.GetSubscribeManager().RegisterSubscribe(NewDeviceAvailabilitySubscribe(adapter.GetAvailabilityTopic("+")))
}

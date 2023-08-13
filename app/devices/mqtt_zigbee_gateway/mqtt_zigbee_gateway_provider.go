package mqtt_zigbee_gateway

import (
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	"github.com/titrxw/smart-home-server/app/mqtt"
	"github.com/titrxw/smart-home-server/app/mqtt/subscribe"
	"github.com/titrxw/smart-home-server/app/pkg/device"
	"strings"
)

type Provider struct {
}

func (p Provider) Register(server *mqtt.Server) {
	config := device.Device{
		Type:              device.DeviceGatewayAppType,
		TypeName:          "mqtt_zigbee_gateway",
		Name:              "mqtt_zigbee网关",
		NeedGateway:       false,
		SupportOperate:    nil,
		OperateDesc:       nil,
		SupportReport:     nil,
		Setting:           nil,
		CtrlTopic:         "/zigbee2mqtt/{app_name}/{appid}/device/{component_appid}/set",
		ReportTopic:       "/zigbee2mqtt/{app_name}/{appid}/device/+/get",
		AvailabilityTopic: "/zigbee2mqtt/{app_name}/{appid}/device/+/availability",
	}
	manager.RegisterDevice(config)

	server.GetSubscribeManager().RegisterSubscribe(subscribe.NewDeviceReportSubscribe(strings.Replace(config.ReportTopic, "{appid}", "+", 1)))
	server.GetSubscribeManager().RegisterSubscribe(NewDeviceAvailabilitySubscribe(strings.Replace(config.AvailabilityTopic, "{appid}", "+", 1)))
}

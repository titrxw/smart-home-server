package mqtt_zigbee_gateway

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	"github.com/titrxw/smart-home-server/app/device_manager/mqtt/subscribe"
)

type DeviceStatus struct {
	State string `json:"state"`
}

type DeviceAvailabilitySubscribe struct {
	subscribe.Abstract
}

func NewDeviceAvailabilitySubscribe(topic string) DeviceAvailabilitySubscribe {
	return DeviceAvailabilitySubscribe{
		subscribe.Abstract{
			Topic:       topic,
			TopicRegexp: subscribe.Abstract{}.MakeTopicRegexpFromTopic(topic),
		},
	}
}

func (s DeviceAvailabilitySubscribe) GetTopic() string {
	return s.GetShareTopic(s.Topic)
}

func (s DeviceAvailabilitySubscribe) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	deviceId := s.GetComponentDeviceIdFromTopic(message.Topic())
	if deviceId != "" {
		deviceStatus := DeviceStatus{}
		err := json.Unmarshal(message.Payload(), &deviceStatus)
		if err == nil {
			device := logic.Logic.Device.GetDeviceByDeviceId(deviceId)
			if device != nil {
				err = logic.Logic.Device.OnOnlineStatucChange(device, "", deviceStatus.State == "online")
			}
		}

		if err != nil {
			helper.ErrLog(err)
		}
	}

}

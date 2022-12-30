package mqtt_zigbee

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	"github.com/titrxw/smart-home-server/app/Mqtt/Subscribe"
)

type DeviceStatus struct {
	State string `json:"state"`
}

type DeviceAvailabilitySubscribe struct {
	subscribe.SubscribeAbstract
}

func NewDeviceAvailabilitySubscribe(topic string) DeviceAvailabilitySubscribe {
	return DeviceAvailabilitySubscribe{
		subscribe.SubscribeAbstract{
			Topic:       topic,
			TopicRegexp: subscribe.SubscribeAbstract{}.MakeTopicRegexpFromTopic(topic),
		},
	}
}

func (deviceAvailabilitySubscribe DeviceAvailabilitySubscribe) GetTopic() string {
	return deviceAvailabilitySubscribe.GetShareTopic(deviceAvailabilitySubscribe.Topic)
}

func (deviceAvailabilitySubscribe DeviceAvailabilitySubscribe) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	deviceId := deviceAvailabilitySubscribe.GetComponentDeviceIdFromTopic(message.Topic())
	if deviceId != "" {
		deviceStatus := DeviceStatus{}
		err := json.Unmarshal(message.Payload(), &deviceStatus)
		if err == nil {
			device := logic.Logic.DeviceLogic.GetDeviceByDeviceId(deviceId)
			if device != nil {
				err = logic.Logic.DeviceLogic.OnOnlineStatucChange(device, "", deviceStatus.State == "online")
			}
		}

		if err != nil {
			global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
		}
	}

}

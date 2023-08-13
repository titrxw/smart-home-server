package mqtt_zigbee_gateway

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/titrxw/smart-home-server/app/mqtt/subscribe"
	devicepkg "github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
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
			iotMessage := &devicepkg.OperateMessage{
				EventType: "device_status_change",
				Payload: map[string]interface{}{
					"status": deviceStatus.State,
				},
			}
			err = logic.Logic.Message.PubClientReportMsg(context.Background(), deviceId, deviceId, iotMessage)
		}

		if err != nil {
			helper.ErrLog(err)
		}
	}

}

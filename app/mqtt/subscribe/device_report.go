package subscribe

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
)

type DeviceReport struct {
	Abstract
}

func NewDeviceReportSubscribe(topic string) DeviceReport {
	return DeviceReport{
		Abstract{
			Topic:       topic,
			TopicRegexp: Abstract{}.MakeTopicRegexpFromTopic(topic),
		},
	}
}

func (s DeviceReport) GetTopic() string {
	return s.GetShareTopic(s.Topic)
}

func (s DeviceReport) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	iotMessage, gatewayDeviceId, deviceId, err := s.validateAndGetPayload(message)
	if err == nil {
		err = logic.Logic.Message.PubClientReportMsg(context.Background(), gatewayDeviceId, deviceId, iotMessage)
	}

	if err != nil {
		helper.ErrLog(err)
	}
}

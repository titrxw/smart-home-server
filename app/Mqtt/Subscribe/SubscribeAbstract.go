package subscribe

import (
	"errors"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/app/Mqtt/Interface"
	"regexp"
	"strings"
)

type SubscribeAbstract struct {
	Interface.SubscribeInterface

	Topic       string
	TopicRegexp string
}

func (subscribe SubscribeAbstract) makeTopicRegexpFromTopic(topic string) string {
	topic = strings.Replace(topic, "/", "\\/", -1)
	topic = strings.Replace(topic, "+", "(.*)", -1)
	topic = "^" + topic

	return topic
}

func (subscribe SubscribeAbstract) getDeviceIdFromTopic(topic string) string {
	reg := regexp.MustCompile(subscribe.TopicRegexp)
	data := reg.FindStringSubmatch(topic)

	if data != nil && data[1] != "" {
		return data[1]
	}

	return ""
}

func (subscribe SubscribeAbstract) validateAndGetPayload(message mqtt.Message) (*cloudevents.Event, *model.Device, error) {
	deviceId := subscribe.getDeviceIdFromTopic(message.Topic())
	if deviceId == "" {
		return nil, nil, errors.New("订阅获取到topic的数据非法")
	}

	device := logic.Logic.DeviceLogic.GetDeviceByDeviceId(deviceId)
	if device == nil {
		return nil, nil, errors.New("订阅获取到topic的数据非法")
	}

	newEvent, err := logic.Logic.EmqxLogic.UnPackMessage(device, string(message.Payload()))
	if err != nil {
		return nil, nil, err
	}

	if newEvent.Source() == "" || newEvent.Type() == "" || newEvent.Subject() == "" {
		return nil, nil, errors.New("订阅获取到payload的数据非法")
	}

	return newEvent, device, nil
}

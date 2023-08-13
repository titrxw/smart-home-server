package subscribe

import (
	deviceInterface "github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"regexp"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Abstract struct {
	Topic       string
	TopicRegexp string
}

// 共享订阅是在多个订阅者之间实现负载均衡的订阅方式 https://www.emqx.io/docs/zh/v5.0/advanced/shared-subscriptions.html#%E5%B8%A6%E7%BE%A4%E7%BB%84%E7%9A%84%E5%85%B1%E4%BA%AB%E8%AE%A2%E9%98%85
func (s Abstract) GetShareTopic(topic string) string {
	return topic
}

func (s Abstract) MakeTopicRegexpFromTopic(topic string) string {
	topic = strings.Replace(topic, "/", "\\/", -1)
	topic = strings.Replace(topic, "+", "(.*)", -1)
	topic = "^" + topic

	return topic
}

func (s Abstract) GetDeviceIdFromTopic(topic string) string {
	reg := regexp.MustCompile(s.TopicRegexp)
	data := reg.FindStringSubmatch(topic)

	if data != nil && data[1] != "" {
		return data[1]
	}

	return ""
}

func (s Abstract) GetComponentDeviceIdFromTopic(topic string) string {
	reg := regexp.MustCompile(s.TopicRegexp)
	data := reg.FindStringSubmatch(topic)

	if data != nil && len(data) == 3 && data[2] != "" {
		return data[2]
	}

	return ""
}

func (s Abstract) validateAndGetPayload(message mqtt.Message) (iotMessage *deviceInterface.OperateMessage, gatewayDeviceId string, deviceId string, err error) {
	gatewayDeviceId = s.GetDeviceIdFromTopic(message.Topic())
	deviceId = s.GetComponentDeviceIdFromTopic(message.Topic())
	if deviceId == "" {
		deviceId = gatewayDeviceId
	}
	if deviceId == "" {
		return nil, "", "", exception.NewRuntimeError("订阅获取到topic的数据非法")
	}

	iotMessage, err = deviceInterface.UnPackMessage(string(message.Payload()))
	if err != nil {
		return nil, "", "", err
	}
	if iotMessage.EventType == "" || iotMessage.Payload == nil {
		return nil, "", "", exception.NewRuntimeError("订阅获取到payload的数据非法")
	}

	return iotMessage, gatewayDeviceId, deviceId, nil
}

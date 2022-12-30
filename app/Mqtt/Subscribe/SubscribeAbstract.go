package subscribe

import (
	"regexp"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/app/Mqtt/Interface"
)

type SubscribeAbstract struct {
	Interface.SubscribeInterface

	Topic       string
	TopicRegexp string
}

//共享订阅是在多个订阅者之间实现负载均衡的订阅方式 https://www.emqx.io/docs/zh/v5.0/advanced/shared-subscriptions.html#%E5%B8%A6%E7%BE%A4%E7%BB%84%E7%9A%84%E5%85%B1%E4%BA%AB%E8%AE%A2%E9%98%85
func (subscribe SubscribeAbstract) GetShareTopic(topic string) string {
	return "$share" + topic
}

func (subscribe SubscribeAbstract) MakeTopicRegexpFromTopic(topic string) string {
	topic = strings.Replace(topic, "/", "\\/", -1)
	topic = strings.Replace(topic, "+", "(.*)", -1)
	topic = "^" + topic

	return topic
}

func (subscribe SubscribeAbstract) GetDeviceIdFromTopic(topic string) string {
	reg := regexp.MustCompile(subscribe.TopicRegexp)
	data := reg.FindStringSubmatch(topic)

	if data != nil && data[1] != "" {
		return data[1]
	}

	return ""
}

func (subscribe SubscribeAbstract) GetComponentDeviceIdFromTopic(topic string) string {
	reg := regexp.MustCompile(subscribe.TopicRegexp)
	data := reg.FindStringSubmatch(topic)

	if data != nil && data[2] != "" {
		return data[2]
	}

	return ""
}

func (subscribe SubscribeAbstract) validateAndGetPayload(message mqtt.Message) (*model.IotMessage, *model.Device, *model.Device, error) {
	deviceId := subscribe.GetDeviceIdFromTopic(message.Topic())
	if deviceId == "" {
		return nil, nil, nil, exception.NewRuntimeError("订阅获取到topic的数据非法")
	}

	gatewayDevice := logic.Logic.DeviceLogic.GetDeviceByDeviceId(deviceId)
	if gatewayDevice == nil {
		return nil, nil, nil, exception.NewRuntimeError("订阅获取到topic的数据非法")
	}
	device := gatewayDevice
	if gatewayDevice.IsGateway() {
		componentAppid := subscribe.GetComponentDeviceIdFromTopic(message.Topic())
		if componentAppid == "" {
			return nil, nil, nil, exception.NewRuntimeError("订阅获取到component_appid的数据非法")
		}
		if componentAppid != gatewayDevice.App.AppId {
			device = logic.Logic.DeviceLogic.GetDeviceByDeviceId(componentAppid)
			if device == nil {
				return nil, nil, nil, exception.NewRuntimeError("订阅获取到topic的数据非法")
			}
		}
	}

	iotMessage, err := logic.Logic.DeviceLogic.GetDeviceAdapter(device.TypeName).UnPackMessage(string(message.Payload()))
	if err != nil {
		return nil, nil, nil, err
	}
	if iotMessage.EventType == "" || iotMessage.Payload == nil {
		return nil, nil, nil, exception.NewRuntimeError("订阅获取到payload的数据非法")
	}

	return iotMessage, gatewayDevice, device, nil
}

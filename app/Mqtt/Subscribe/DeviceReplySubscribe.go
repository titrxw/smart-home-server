package subscribe

import (
	"reflect"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type DeviceReplaySubscribe struct {
	SubscribeAbstract
}

func NewDeviceReplaySubscribe(topic string) DeviceReplaySubscribe {
	return DeviceReplaySubscribe{
		SubscribeAbstract{
			Topic:       topic,
			TopicRegexp: SubscribeAbstract{}.makeTopicRegexpFromTopic(topic),
		},
	}
}

func (deviceReplaySubscribe DeviceReplaySubscribe) GetTopic() string {
	return deviceReplaySubscribe.Topic
}

func (deviceReplaySubscribe DeviceReplaySubscribe) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	cloudEvent, device, err := deviceReplaySubscribe.validateAndGetPayload(message)
	if err == nil {
		operateLog, err := logic.Logic.DeviceOperateLogic.OnOperateResponse(device, cloudEvent)
		if err == nil {
			err = logic.Logic.DeviceLogic.GetDeviceAdapter(operateLog.DeviceType).OnOperateResponse(client, device, operateLog, cloudEvent)
			global.FApp.Event.Publish(reflect.TypeOf(event.DeviceOperateReplyEvent{}).Name(), event.NewDeviceOperateReplyEvent(device, operateLog, cloudEvent))
		}
	}

	if err != nil {
		global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
	}
}

package subscribe

import (
	"encoding/json"
	"reflect"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
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
	cloudEvent, err := deviceReplaySubscribe.validateAndGetPayload(message)
	if err == nil {
		replyMessage := model.DeviceOperateReplyMessage{}
		err := json.Unmarshal(cloudEvent.Data(), &replyMessage)
		if err == nil {
			device, operateLog, err := logic.Logic.DeviceOperateLogic.OnOperateResponse(replyMessage)
			if err == nil {
				err = logic.Logic.DeviceLogic.GetDeviceAdapter(operateLog.Type).OnOperateResponse(device, operateLog, cloudEvent)
				global.FApp.Event.Publish(reflect.TypeOf(event.DeviceOperateReplyEvent{}).Name(), event.NewDeviceOperateReplyEvent(device, operateLog, cloudEvent))
			}
		}

	}

	if err != nil {
		global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
	}
}

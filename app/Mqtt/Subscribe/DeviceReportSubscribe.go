package subscribe

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	"reflect"
)

type DeviceReportSubscribe struct {
	SubscribeAbstract
}

func NewDeviceReportSubscribe(topic string) DeviceReportSubscribe {
	return DeviceReportSubscribe{
		SubscribeAbstract{
			Topic:       topic,
			TopicRegexp: SubscribeAbstract{}.makeTopicRegexpFromTopic(topic),
		},
	}
}

func (deviceReportSubscribe DeviceReportSubscribe) GetTopic() string {
	return deviceReportSubscribe.Topic
}

func (deviceReportSubscribe DeviceReportSubscribe) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	cloudEvent, err := deviceReportSubscribe.validateAndGetPayload(message)
	if err == nil {
		device := logic.Logic.DeviceLogic.GetDeviceByDeviceId(cloudEvent.ID())
		if device != nil {
			err := logic.Logic.DeviceLogic.GetDeviceAdapter(device.Type).OnReport(device, cloudEvent)
			global.FApp.Event.Publish(reflect.TypeOf(event.DeviceReportEvent{}).Name(), event.NewDeviceReportEvent(device, cloudEvent))

			if err != nil {
				global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
			}
		}
	}
}

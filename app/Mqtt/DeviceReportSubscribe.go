package mqtt

import (
	"reflect"
	"regexp"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type DeviceReportSubscribe struct {
	topic string
}

func NewDeviceReportSubscribe(topic string) DeviceReportSubscribe {
	topic = strings.Replace(topic, "/", "\\/", -1)
	topic = strings.Replace(topic, "+", "(.*)", -1)
	topic = "^" + topic

	return DeviceReportSubscribe{
		topic: topic,
	}
}

func (deviceReportSubscribe DeviceReportSubscribe) DeviceReportSubscribe(client mqtt.Client, message mqtt.Message) {
	reg := regexp.MustCompile(deviceReportSubscribe.topic)
	data := reg.FindStringSubmatch(message.Topic())
	if data[1] != "" {
		device := logic.Logic.DeviceLogic.GetDeviceByDeviceId(data[1])
		if device != nil {
			payload := string(message.Payload())
			err := logic.Logic.DeviceLogic.GetDeviceAdapter(device.Type).OnReport(device, payload)
			global.FApp.Event.Publish(reflect.TypeOf(event.DeviceReportEvent{}).Name(), event.NewDeviceReportEvent(device, payload))

			if err != nil {
				global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
			}
		}
	}
}

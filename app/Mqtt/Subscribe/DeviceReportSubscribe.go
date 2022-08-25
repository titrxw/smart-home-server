package subscribe

import (
	"reflect"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

const CTRL_EXCEPTION = "ctrl_exception"

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
	cloudEvent, device, err := deviceReportSubscribe.validateAndGetPayload(message)
	if err == nil {
		if cloudEvent.Type() == CTRL_EXCEPTION {
			_, err = logic.Logic.DeviceOperateLogic.OnOperateResponse(device, cloudEvent)
		}

		if err == nil {
			if cloudEvent.ID() == device.App.AppId {
				cloudEvent.SetID(logic.Logic.DeviceReportLogic.GetOperateOrReportNumber(device.App.AppId))
			}
			reportLog, err := logic.Logic.DeviceReportLogic.OnReport(device, cloudEvent)
			if err == nil {
				err = logic.Logic.DeviceLogic.GetDeviceAdapter(device.Type).OnReport(client, device, reportLog, cloudEvent)
				global.FApp.Event.Publish(reflect.TypeOf(event.DeviceReportEvent{}).Name(), event.NewDeviceReportEvent(device, cloudEvent))
			}
		}

		if err != nil {
			global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
		}
	}
}

package subscribe

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type DeviceReportSubscribe struct {
	SubscribeAbstract
}

func NewDeviceReportSubscribe(topic string) DeviceReportSubscribe {
	return DeviceReportSubscribe{
		SubscribeAbstract{
			Topic:       topic,
			TopicRegexp: SubscribeAbstract{}.MakeTopicRegexpFromTopic(topic),
		},
	}
}

func (deviceReportSubscribe DeviceReportSubscribe) GetTopic() string {
	return deviceReportSubscribe.GetShareTopic(deviceReportSubscribe.Topic)
}

func (deviceReportSubscribe DeviceReportSubscribe) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	iotMessage, gatewayDevice, device, err := deviceReportSubscribe.validateAndGetPayload(message)
	if err == nil {
		if iotMessage.Id != "" {
			err = logic.Logic.DeviceOperateLogic.OnOperateResponse(gatewayDevice, device, iotMessage)
		} else {
			err = logic.Logic.DeviceReportLogic.OnReport(gatewayDevice, device, iotMessage)
		}
	}

	if err != nil {
		global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
	}
}

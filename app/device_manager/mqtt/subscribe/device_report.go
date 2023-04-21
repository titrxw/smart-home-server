package subscribe

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
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
	iotMessage, gatewayDevice, device, err := s.validateAndGetPayload(message)
	if err == nil {
		if iotMessage.Id != "" {
			err = logic.Logic.DeviceOperate.OnOperateResponse(gatewayDevice, device, iotMessage)
		} else {
			err = logic.Logic.DeviceReport.OnReport(gatewayDevice, device, iotMessage)
		}
	}

	if err != nil {
		helper.ErrLog(err)
	}
}

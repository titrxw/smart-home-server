package subscribe

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	"regexp"
)

type DeviceStatus struct {
	UserName string `json:"username"`
	Ip       string `json:"ipaddress"`
}

type DeviceStatusChangeSubscribe struct {
	SubscribeAbstract
}

func NewDeviceStatusChangeSubscribe(topic string) DeviceStatusChangeSubscribe {
	return DeviceStatusChangeSubscribe{
		SubscribeAbstract{
			Topic:       topic,
			TopicRegexp: "",
		},
	}
}

func (deviceStatusChangeSubscribe DeviceStatusChangeSubscribe) GetTopic() string {
	return deviceStatusChangeSubscribe.Topic
}

func (deviceStatusChangeSubscribe DeviceStatusChangeSubscribe) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	//"^\$SYS\/brokers\/.*\/clients\/.*\/(dis)?connected"
	//"$SYS/brokers/+/clients/+/+"
	reg := regexp.MustCompile(`^\$SYS\/brokers\/.*\/clients\/(.*)\/(.*)`)
	data := reg.FindStringSubmatch(message.Topic())
	if data[2] == "connected" || data[2] == "disconnected" {
		deviceStatus := DeviceStatus{}
		err := json.Unmarshal(message.Payload(), &deviceStatus)
		if err == nil {
			device := logic.Logic.DeviceLogic.GetDeviceByDeviceId(deviceStatus.UserName)
			if device != nil {
				err = logic.Logic.DeviceLogic.OnOnlineStatucChange(device, deviceStatus.Ip, data[2] == "connected")
			}
		}

		if err != nil {
			global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
		}
	}
}

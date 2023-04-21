package subscribe

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	"regexp"
)

type DeviceStatus struct {
	UserName string `json:"username"`
	Ip       string `json:"ipaddress"`
}

type DeviceStatusChange struct {
	Abstract
}

func NewDeviceStatusChangeSubscribe(topic string) DeviceStatusChange {
	return DeviceStatusChange{
		Abstract{
			Topic:       topic,
			TopicRegexp: "",
		},
	}
}

func (s DeviceStatusChange) GetTopic() string {
	return s.Topic
}

func (s DeviceStatusChange) OnSubscribe(client mqtt.Client, message mqtt.Message) {
	//"^\$SYS\/brokers\/.*\/clients\/.*\/(dis)?connected"
	//"$SYS/brokers/+/clients/+/+"
	reg := regexp.MustCompile(`^\$SYS\/brokers\/.*\/clients\/(.*)\/(.*)`)
	data := reg.FindStringSubmatch(message.Topic())
	if data[2] == "connected" || data[2] == "disconnected" {
		deviceStatus := DeviceStatus{}
		err := json.Unmarshal(message.Payload(), &deviceStatus)
		if err == nil {
			device := logic.Logic.Device.GetDeviceByDeviceId(deviceStatus.UserName)
			if device != nil {
				err = logic.Logic.Device.OnOnlineStatucChange(device, deviceStatus.Ip, data[2] == "connected")
			}
		}

		if err != nil {
			helper.ErrLog(err)
		}
	}
}

package subscribe

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	devicepkg "github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
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
			state := "offline"
			if data[2] == "connected" {
				state = "online"
			}
			iotMessage := &devicepkg.OperateMessage{
				EventType: "device_status_change",
				Payload: map[string]interface{}{
					"status": state,
				},
			}
			err = logic.Logic.Message.PubClientReportMsg(context.Background(), deviceStatus.UserName, deviceStatus.UserName, iotMessage)
		}

		if err != nil {
			helper.ErrLog(err)
		}
	}
}

package mqtt

import (
	"context"
	"encoding/json"
	"regexp"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceStatus struct {
	UserName string `json:"username"`
	Ip       string `json:"ipaddress"`
}

type DeviceOffline struct {
	UserName string `json:"username"`
}

type DeviceStatusChangeSubscribe struct {
}

func (deviceStatusChangeSubscribe DeviceStatusChangeSubscribe) DeviceStatusChangeSubscribe(client mqtt.Client, message mqtt.Message) {
	//"^\$SYS\/brokers\/.*\/clients\/.*\/(dis)?connected"
	reg := regexp.MustCompile(`^\$SYS\/brokers\/.*\/clients\/(.*)\/(.*)`)
	data := reg.FindStringSubmatch(message.Topic())
	if data[2] == "connected" || data[2] == "disconnected" {
		deviceStatus := DeviceStatus{}
		err := json.Unmarshal(message.Payload(), &deviceStatus)
		if err == nil {
			device := logic.Logic.DeviceLogic.GetDeviceByDeviceId(deviceStatus.UserName)
			if device != nil {
				if data[2] == "connected" {
					device.OnlineStatus = model.DEVICE_ONLINE
					device.LastIp = deviceStatus.Ip
					device.LatestVisit = time.Now().Format(model.TimeFormat)
				} else {
					device.OnlineStatus = model.DEVICE_OFFLINE
				}
				err = logic.Logic.DeviceLogic.UpdateDevice(context.Background(), device)
			}
		}

		if err != nil {
			global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
		}
	}
}

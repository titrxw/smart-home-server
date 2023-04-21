package _interface

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Interface interface {
	GetTopic() string
	OnSubscribe(client mqtt.Client, message mqtt.Message)
}

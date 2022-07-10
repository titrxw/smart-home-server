package Interface

import mqtt "github.com/eclipse/paho.mqtt.golang"

type SubscribeInterface interface {
	GetTopic() string
	OnSubscribe(client mqtt.Client, message mqtt.Message)
}

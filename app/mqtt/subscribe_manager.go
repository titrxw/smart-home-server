package mqtt

import (
	"context"
	"github.com/titrxw/smart-home-server/app/common/emqx"
	"github.com/titrxw/smart-home-server/app/common/helper"
	subscribeInterface "github.com/titrxw/smart-home-server/app/mqtt/interface"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SubscribeManager struct {
	subscribeMap map[string]mqtt.MessageHandler
}

var subscribeManager *SubscribeManager

func GetSubscribeManager() *SubscribeManager {
	if subscribeManager == nil {
		subscribeManager = &SubscribeManager{
			subscribeMap: make(map[string]mqtt.MessageHandler),
		}
	}

	return subscribeManager
}

func (sm *SubscribeManager) RegisterSubscribe(subscribeInterface subscribeInterface.Interface) {
	subscribeManager.subscribeMap[subscribeInterface.GetTopic()] = subscribeInterface.OnSubscribe
}

func (sm *SubscribeManager) Start(EMQServerAddress string, port string, userName string, password string) {
	sm.initMqttDevice(userName, password)
	sm.startMqttClient(EMQServerAddress, port, userName, password)
}

func (sm *SubscribeManager) initMqttDevice(userName string, password string) {
	ctx := context.Background()
	emqService := emqx.GetEmqxClientService()
	emqService.DeleteClient(ctx, userName)
	err := emqService.AddClient(ctx, userName, password, "")
	if err != nil {
		panic(err)
	}

	for topic, _ := range sm.subscribeMap {
		err = emqService.AddClientSubAcl(context.Background(), userName, topic)
		if err != nil {
			panic(err)
		}
	}
}

func (sm *SubscribeManager) startMqttClient(EMQServerAddress string, port string, userName string, password string) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + EMQServerAddress + ":" + port)
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(60 * time.Second)
	opts.SetClientID(userName)
	opts.SetCleanSession(false)
	opts.SetAutoReconnect(true)
	opts.OnConnect = func(client mqtt.Client) {
		for topic, callback := range sm.subscribeMap {
			sm.subscribe(client, topic, callback)
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (sm *SubscribeManager) subscribe(client mqtt.Client, topic string, callback mqtt.MessageHandler) {
	token := client.Subscribe(topic, 2, callback)
	if token.Wait() && token.Error() != nil {
		helper.ErrLog(token.Error())
	}
}

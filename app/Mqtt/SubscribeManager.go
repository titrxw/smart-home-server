package mqtt

import (
	"context"
	subscribeInterface "github.com/titrxw/smart-home-server/app/Mqtt/Interface"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	emqx "github.com/titrxw/smart-home-server/app/Service/Emqx"
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

func (subscribeManager *SubscribeManager) RegisterSubscribe(subscribeInterface subscribeInterface.SubscribeInterface) {
	subscribeManager.subscribeMap[subscribeInterface.GetTopic()] = subscribeInterface.OnSubscribe
}

func (subscribeManager SubscribeManager) Start(EMQServerAddress string, port string, userName string, password string) {
	subscribeManager.initMqttDevice(userName, password)
	subscribeManager.startMqttClient(EMQServerAddress, port, userName, password)
}

func (subscribeManager SubscribeManager) initMqttDevice(userName string, password string) {
	ctx := context.Background()
	emqService := emqx.GetEmqxClientService(global.FApp.Container)
	emqService.DeleteClient(ctx, userName)
	err := emqService.AddClient(ctx, userName, password, "")
	if err != nil {
		panic(err)
	}

	for topic, _ := range subscribeManager.subscribeMap {
		err = emqService.AddClientSubAcl(context.Background(), userName, topic)
		if err != nil {
			panic(err)
		}
	}
}

func (subscribeManager SubscribeManager) startMqttClient(EMQServerAddress string, port string, userName string, password string) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + EMQServerAddress + ":" + port)
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(60 * time.Second)
	opts.SetClientID(userName)
	opts.SetCleanSession(false)
	opts.SetAutoReconnect(true)
	opts.OnConnect = func(client mqtt.Client) {
		for topic, callback := range subscribeManager.subscribeMap {
			subscribeManager.subscribe(client, topic, callback)
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (subscribeManager SubscribeManager) subscribe(client mqtt.Client, topic string, callback mqtt.MessageHandler) {
	token := client.Subscribe(topic, 2, callback)
	if token.Wait() && token.Error() != nil {
		global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, token.Error(), "")
	}
}

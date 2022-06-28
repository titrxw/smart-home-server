package mqtt

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	emqx "github.com/titrxw/smart-home-server/app/Service/Emqx"
	"reflect"
	"time"
)

type DeviceReplaySubscribe struct {
}

func initDevice(userName string, password string) {
	ctx := context.Background()
	emqService := emqx.GetEmqxClientService(global.FApp.Container)
	emqService.DeleteClient(ctx, userName)
	err := emqService.AddClient(ctx, userName, password, "")
	if err != nil {
		panic(err)
	}
	err = emqService.AddClientSubAcl(context.Background(), userName, logic.Logic.GetClientPubTopic("+"))
	if err != nil {
		panic(err)
	}
}

func StartDeviceReplaySubscribe(EMQServerAddress string, port string, userName string, password string) {
	initDevice(userName, password)

	opts := mqtt.NewClientOptions().AddBroker("tcp://" + EMQServerAddress + ":" + port)
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(60 * time.Second)
	opts.SetClientID(userName)
	opts.SetCleanSession(false)
	opts.SetAutoReconnect(true)
	opts.OnConnect = func(client mqtt.Client) {
		DeviceReplaySubscribe{}.Subscribe(client, logic.Logic.EmqxLogic.GetClientPubTopic("+"))
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (this DeviceReplaySubscribe) onSubscribe(client mqtt.Client, message mqtt.Message) {
	replyMessage := model.DeviceOperateReplyMessage{}
	err := json.Unmarshal(message.Payload(), &replyMessage)
	if err == nil {
		operateLog, err := logic.Logic.DeviceOperateLogic.GetOperateLogResultByNumber(replyMessage.OperateId)
		if err == nil {
			operateLog.ResponsePayload = replyMessage.PayLoad
			operateLog.ResponseTime = time.Now().Format(model.TimeFormat)
			err := logic.Logic.DeviceOperateLogic.UpdateOperateLog(operateLog)
			if err == nil {
				err = logic.Logic.DeviceLogic.GetDeviceAdapter(operateLog.Type).OnOperateResponse(operateLog)
				global.FApp.Event.Publish(reflect.TypeOf(event.DeviceOperateReplyEvent{}).Name(), event.NewDeviceOperateReplyEvent(operateLog))
			}
		}
	}

	if err != nil {
		global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, err, "")
	}
}

func (this DeviceReplaySubscribe) Subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 2, this.onSubscribe)
	if token.Wait() && token.Error() != nil {
		global.FApp.HandlerExceptions.GetExceptionHandler().Reporter(global.FApp.HandlerExceptions.Logger, token.Error(), "")
	}
}

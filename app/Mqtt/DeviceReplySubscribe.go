package mqtt

import (
	"encoding/json"
	"reflect"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	global "github.com/titrxw/go-framework/src/Global"
	event "github.com/titrxw/smart-home-server/app/Event"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceReplaySubscribe struct {
}

func (deviceReplaySubscribe DeviceReplaySubscribe) DeviceReplySubscribe(client mqtt.Client, message mqtt.Message) {
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

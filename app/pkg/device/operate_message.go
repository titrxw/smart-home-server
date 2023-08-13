package device

import "github.com/titrxw/smart-home-server/app/pkg/helper"

type OperateMessage struct {
	Id        string                 `json:"id"`
	EventType string                 `json:"event_type"`
	Payload   map[string]interface{} `json:"payload"`
	Timestamp int64                  `json:"timestamp"`
}

func PackMessage(message *OperateMessage) (string, error) {
	return helper.JsonEncode(message)
}

func UnPackMessage(payload string) (*OperateMessage, error) {
	message := new(OperateMessage)
	return message, helper.JsonDecode(payload, message)
}

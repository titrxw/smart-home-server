package _interface

import (
	"github.com/titrxw/smart-home-server/app/common/helper"
)

type Abstract struct {
	DeviceInterface
}

func (a Abstract) GetCtrlTopic(appId string, componentAppId string) string {
	return "/iot/" + helper.GetAppName() + "/device/" + appId + "/ctrl"
}

func (a Abstract) GetReportTopic(appId string) string {
	return "/iot/" + helper.GetAppName() + "/device/" + appId + "/report"
}

func (a Abstract) GetAvailabilityTopic(appId string) string {
	return ""
}

func (a Abstract) PackMessage(message *DeviceOperateMessage) (string, error) {
	return helper.JsonEncode(message)
}

func (a Abstract) UnPackMessage(payload string) (*DeviceOperateMessage, error) {
	message := new(DeviceOperateMessage)
	return message, helper.JsonDecode(payload, message)
}

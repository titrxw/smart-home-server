package Interface

import (
	global "github.com/titrxw/go-framework/src/Global"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceAdapterAbstract struct {
	DeviceAdapterInterface
}

func (abstract DeviceAdapterAbstract) GetCtrlTopic(appId string, componentAppId string) string {
	return "/iot/" + global.FApp.Name + "/device/" + appId + "/ctrl"
}

func (faceIdentifyDeviceAdapter DeviceAdapterAbstract) GetReportTopic(appId string) string {
	return "/iot/" + global.FApp.Name + "/device/" + appId + "/report"
}

func (faceIdentifyDeviceAdapter DeviceAdapterAbstract) GetAvailabilityTopic(appId string) string {
	return ""
}

func (faceIdentifyDeviceAdapter DeviceAdapterAbstract) PackMessage(message *model.IotMessage) (string, error) {
	return helper.JsonEncode(message)
}

func (faceIdentifyDeviceAdapter DeviceAdapterAbstract) UnPackMessage(payload string) (*model.IotMessage, error) {
	message := new(model.IotMessage)
	return message, helper.JsonDecode(payload, message)
}

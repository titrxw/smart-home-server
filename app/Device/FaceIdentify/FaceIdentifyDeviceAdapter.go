package faceIdentify

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
)

type FaceIdentifyDeviceAdapter struct {
	Interface.DeviceAdapterInterface
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) GetDeviceConfig() config.Device {
	return config.Device{
		Type:           "face_identify",
		Name:           "识别",
		SupportOperate: []string{"train", "identify"},
		OperateDesc:    map[string]string{"train": "训练", "identify": "识别"},
		SupportReport:  []string{"identify"},
	}
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) OnOperateResponse(device *model.Device, deviceOperateLog *model.DeviceOperateLog, cloudEvent *cloudevents.Event) error {
	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) OnReport(device *model.Device, deviceReportLog *model.DeviceReportLog, cloudEvent *cloudevents.Event) error {
	return nil
}

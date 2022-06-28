package Interface

import model "github.com/titrxw/smart-home-server/app/Model"

type DeviceAdapterInterface interface {
	GetDeviceType() string
	BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error
	OnOperateResponse(deviceOperateLog *model.DeviceOperateLog) error
}

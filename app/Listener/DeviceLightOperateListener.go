package listener

import event "github.com/titrxw/smart-home-server/app/Event"

type DeviceLightOperateListener struct {
}

func (this DeviceLightOperateListener) Handle(event event.DeviceOperateTriggerEvent) {
	if event.Device.Type == "light" {

	}
}

package manager

import (
	"github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"strings"
)

var supportDevice = make(map[string]device.Device)

func RegisterDevice(device device.Device) {
	if device.CtrlTopic == "" {
		device.CtrlTopic = "/iot/{app_name}/device/{appid}/ctrl"
	}
	if device.ReportTopic == "" {
		device.ReportTopic = "/iot/{app_name}/device/{appid}/report"
	}
	device.CtrlTopic = strings.Replace(device.CtrlTopic, "{app_name}", helper.GetAppName(), 1)
	device.ReportTopic = strings.Replace(device.ReportTopic, "{app_name}", helper.GetAppName(), 1)
	device.AvailabilityTopic = strings.Replace(device.AvailabilityTopic, "{app_name}", helper.GetAppName(), 1)
	supportDevice[device.TypeName] = device
}

func GetDeviceSupportMap() map[string]device.Device {
	return supportDevice
}

func GetDeviceConfigByDeviceType(deviceType string) device.Device {
	return GetDeviceSupportMap()[deviceType]
}

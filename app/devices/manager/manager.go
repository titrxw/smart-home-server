package manager

import (
	deviceInterface "github.com/titrxw/smart-home-server/app/devices/interface"
)

var supportDevice = make(map[string]deviceInterface.Device)
var supportDeviceAdapter = make(map[string]deviceInterface.DeviceInterface)

func RegisterDevice(adapterInterface deviceInterface.DeviceInterface) {
	deviceConfig := adapterInterface.GetDeviceConfig()
	supportDeviceAdapter[deviceConfig.TypeName] = adapterInterface
	supportDevice[deviceConfig.TypeName] = deviceConfig
}

func GetDeviceSupportMap() map[string]deviceInterface.Device {
	return supportDevice
}

func GetDeviceByDeviceType(deviceType string) deviceInterface.Device {
	return GetDeviceSupportMap()[deviceType]
}

func GetDevice(deviceType string) deviceInterface.DeviceInterface {
	return supportDeviceAdapter[deviceType]
}

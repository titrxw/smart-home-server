package provider

import (
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	device "github.com/titrxw/smart-home-server/app/Adapter/Device"
	"github.com/titrxw/smart-home-server/app/Adapter/Interface"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	"github.com/titrxw/smart-home-server/config"
)

type DeviceProvider struct {
	provider.ProviderAbstract
}

func (deviceProvider *DeviceProvider) Register(options interface{}) {
	deviceProvider.initDeviceMapInfo()
	deviceProvider.registerLight()
}

func (deviceProvider *DeviceProvider) initDeviceMapInfo() {
	logic.Logic.DeviceLogic.SupportDeviceMap = make(map[string]config.Device)
	logic.Logic.DeviceLogic.SupportDeviceAdapter = make(map[string]Interface.DeviceAdapterInterface)
}

func (deviceProvider *DeviceProvider) registerLight() {
	lightDeviceAdapter := new(device.LightDeviceAdapter)
	logic.Logic.DeviceLogic.RegisterDevice(lightDeviceAdapter.GetDeviceType(), config.Device{
		Name:           "电灯",
		SupportOperate: []string{"on", "off"},
		OperateDesc:    map[string]string{"on": "开灯", "off": "关灯"},
	})
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(lightDeviceAdapter.GetDeviceType(), lightDeviceAdapter)
}

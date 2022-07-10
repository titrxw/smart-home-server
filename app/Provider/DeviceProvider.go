package provider

import (
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	light "github.com/titrxw/smart-home-server/app/Device/Light"
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
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(new(light.LightDeviceAdapter))
}

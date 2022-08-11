package light

import (
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type LightDeviceProvider struct {
	provider.ProviderAbstract
}

func (lightDeviceProvider *LightDeviceProvider) Register(options interface{}) {
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(new(LightDeviceAdapter))
}

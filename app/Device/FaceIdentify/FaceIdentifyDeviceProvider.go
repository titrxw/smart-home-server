package faceIdentify

import (
	provider "github.com/titrxw/go-framework/src/Core/Provider"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type FaceIdentifyDeviceProvider struct {
	provider.ProviderAbstract
}

func (faceIdentifyDeviceProvider *FaceIdentifyDeviceProvider) Register(options interface{}) {
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(new(FaceIdentifyDeviceAdapter))
}

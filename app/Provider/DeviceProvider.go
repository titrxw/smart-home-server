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

func (this *DeviceProvider) Register(options interface{}) {
	this.initDeviceMapInfo()
	this.registerLight()
}

func (this *DeviceProvider) initDeviceMapInfo() {
	logic.Logic.DeviceLogic.SupportDeviceMap = make(map[string]config.Device)
	logic.Logic.DeviceLogic.SupportDeviceAdapter = make(map[string]Interface.DeviceAdapterInterface)
}

func (this *DeviceProvider) registerLight() {
	lightDeviceAdapter := new(device.LightDeviceAdapter)
	logic.Logic.DeviceLogic.RegisterDevice(lightDeviceAdapter.GetDeviceType(), config.Device{
		Name:           "电灯",
		SupportOperate: []string{"on", "off"},
	})
	logic.Logic.DeviceLogic.RegisterDeviceAdapter(lightDeviceAdapter.GetDeviceType(), lightDeviceAdapter)
	//err := global.FApp.Event.Subscribe(reflect.TypeOf(event.DeviceOperateEvent{}).Name(), listener.DeviceLightOperateListener{}.Handle)
	//if err != nil {
	//	panic(err)
	//}
}

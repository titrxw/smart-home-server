package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
)

type DeviceGateway struct {
	Abstract
}

func (l DeviceGateway) UserGatewayAddDevice(ctx context.Context, userId model.UID, gatewayDeviceId uint, deviceId uint) error {
	gatewayDevice, err := Logic.Device.GetUserDeviceById(userId, gatewayDeviceId)
	if err != nil {
		return err
	}
	if !gatewayDevice.IsGateway() {
		return exception.NewResponseError("网关参数错误")
	}
	device, err := Logic.Device.GetUserDeviceById(userId, deviceId)
	if err != nil {
		return err
	}
	if device.IsGateway() {
		return exception.NewResponseError("设备参数错误")
	}
	device.GatewayDeviceId = gatewayDevice.ID
	return Logic.Device.UpdateDevice(ctx, device)
}

func (l DeviceGateway) GetGatewayDevice(ctx context.Context, device *model.Device) *model.Device {
	return Logic.Device.GetDeviceById(device.GatewayDeviceId)
}

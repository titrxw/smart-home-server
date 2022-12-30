package logic

import (
	"context"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type DeviceGatewayLogic struct {
	LogicAbstract
}

func (deviceGatewayLogic DeviceGatewayLogic) UserGatewayAddDevice(ctx context.Context, userId model.UID, gatewayDeviceId uint, deviceId uint) error {
	gatewayDevice, err := Logic.DeviceLogic.GetUserDeviceById(userId, gatewayDeviceId)
	if err != nil {
		return err
	}
	if !gatewayDevice.IsGateway() {
		return exception.NewLogicError("网关参数错误")
	}
	device, err := Logic.DeviceLogic.GetUserDeviceById(userId, deviceId)
	if err != nil {
		return err
	}
	if device.IsGateway() {
		return exception.NewLogicError("设备参数错误")
	}
	device.GatewayDeviceId = gatewayDevice.ID
	return Logic.DeviceLogic.UpdateDevice(ctx, device)
}

func (deviceGatewayLogic DeviceGatewayLogic) GetGatewayDevice(ctx context.Context, device *model.Device) *model.Device {
	return Logic.DeviceLogic.GetDeviceById(device.GatewayDeviceId)
}

package logic

import (
	"context"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type MessageLogic struct {
}

func (msgLogic MessageLogic) PubClientOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, message *model.IotMessage) error {
	adapter := Logic.DeviceLogic.GetDeviceAdapter(device.TypeName)
	payload, err := adapter.PackMessage(message)
	if err != nil {
		return err
	}

	return Logic.EmqxLogic.PubClientOperate(ctx, gatewayDevice.App.AppId, Logic.DeviceLogic.GetDeviceAdapter(gatewayDevice.TypeName).GetCtrlTopic(gatewayDevice.App.AppId, device.App.AppId), payload, 2)
}

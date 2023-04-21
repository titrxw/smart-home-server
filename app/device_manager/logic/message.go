package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	_interface "github.com/titrxw/smart-home-server/app/devices/interface"
	"github.com/titrxw/smart-home-server/app/devices/manager"
)

type Message struct {
	Abstract
}

func (l Message) PubClientOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, message *_interface.DeviceOperateMessage) error {
	adapter := manager.GetDevice(device.TypeName)
	payload, err := adapter.PackMessage(message)
	if err != nil {
		return err
	}

	return Logic.Emqx.PubClientOperate(ctx, gatewayDevice.App.AppId, manager.GetDevice(gatewayDevice.TypeName).GetCtrlTopic(gatewayDevice.App.AppId, device.App.AppId), payload, 2)
}

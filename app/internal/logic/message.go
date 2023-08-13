package logic

import (
	"context"
	"errors"
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	"github.com/titrxw/smart-home-server/app/internal/model"
	devicepkg "github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
	"strings"
)

type Message struct {
	logic.Abstract
}

func (l *Message) PubClientOperateMsg(ctx context.Context, gatewayDevice *model.Device, device *model.Device, message *devicepkg.OperateMessage) error {
	payload, err := devicepkg.PackMessage(message)
	if err != nil {
		return err
	}

	deviceConfig := manager.GetDeviceConfigByDeviceType(gatewayDevice.TypeName)
	topic := strings.Replace(deviceConfig.CtrlTopic, "{appid}", gatewayDevice.App.AppId, 1)
	topic = strings.Replace(topic, "{component_appid}", device.App.AppId, 1)

	return Logic.Emqx.PubClientOperate(ctx, gatewayDevice.App.AppId, topic, payload, 2)
}

func (l *Message) PubClientReportMsg(ctx context.Context, gatewayDeviceId string, deviceId string, message *devicepkg.OperateMessage) error {
	gatewayDevice := Logic.Device.GetDeviceByDeviceId(gatewayDeviceId)
	if gatewayDevice == nil {
		return errors.New("非法请求")
	}

	var device *model.Device
	device = gatewayDevice
	if gatewayDeviceId != deviceId {
		device = Logic.Device.GetDeviceByDeviceId(deviceId)
		if device == nil {
			return errors.New("非法请求")
		}
	}

	if message.EventType == "device_status_change" {
		result, exists := message.Payload["status"]
		if exists {
			status, is := result.(string)
			if is {
				return Logic.Device.OnOnlineStatusChange(device, "", status == "online")
			}
		}
		return errors.New("非法请求")
	}

	var err error
	if message.Id != "" {
		err = Logic.DeviceOperate.OnOperateResponse(ctx, gatewayDevice, device, message)
	} else {
		err = Logic.DeviceReport.OnReport(ctx, gatewayDevice, device, message)
	}

	if err == nil {
		Logic.Device.PushMsgToDeviceServer(ctx, gatewayDevice, device, message)
	}

	return err
}

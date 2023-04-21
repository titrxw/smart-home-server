package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/device_manager/event"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/titrxw/smart-home-server/app/device_manager/repository"
	"github.com/titrxw/smart-home-server/app/devices/manager"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type Device struct {
	Abstract
}

func (l Device) IsSupportOperate(device *model.Device, operate model.DeviceOperateType) bool {
	supportOperateMap := manager.GetDeviceByDeviceType(device.TypeName).SupportOperate
	if supportOperateMap == nil {
		return false
	}

	for _, element := range supportOperateMap {
		if operate == model.DeviceOperateType(element) {
			return true
		}
	}

	return false
}

func (l Device) IsSupportReport(device *model.Device, operate model.DeviceReportType) bool {
	supportReportMap := manager.GetDeviceByDeviceType(device.TypeName).SupportReport
	if supportReportMap == nil {
		return false
	}

	for _, element := range supportReportMap {
		if operate == model.DeviceReportType(element) {
			return true
		}
	}

	return false
}

func (l Device) GetDeviceByDeviceId(deviceId string) *model.Device {
	deviceApp := Logic.App.GetAppByAppId(deviceId)
	if deviceApp == nil {
		return nil
	}

	return repository.Repository.Device.GetDeviceByApp(l.GetDefaultDb(), deviceApp)
}

func (l Device) GetDeviceById(deviceId uint) *model.Device {
	return repository.Repository.Device.GetDeviceById(l.GetDefaultDb(), deviceId)
}

func (l Device) CreateUserDevice(ctx context.Context, userId model.UID, deviceName string, deviceType string) (*model.Device, error) {
	device := &model.Device{
		Name:     deviceName,
		TypeName: deviceType,
		Type:     manager.GetDeviceSupportMap()[deviceType].Type,
	}

	err := l.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		deviceApp := repository.Repository.App.CreateDeviceApp(tx, device.Type)
		if deviceApp == nil {
			return exception.NewResponseError("创建app失败")
		}

		device = repository.Repository.Device.AddUserDeviceByApp(tx, userId, deviceApp, device)
		if device == nil {
			return exception.NewResponseError("创建设备失败")
		}

		if !manager.GetDeviceSupportMap()[deviceType].NeedGateway {
			adapter := manager.GetDevice(deviceType)
			return Logic.Emqx.AddEmqxClient(ctx, device, map[string][]string{
				"pub": {
					adapter.GetReportTopic(deviceApp.AppId),
					adapter.GetAvailabilityTopic(deviceApp.AppId),
				},
				"sub": {
					adapter.GetCtrlTopic(deviceApp.AppId, "+"),
				},
			})
		}

		return nil
	})

	return device, err
}

func (l Device) GetUserDevices(userId model.UID, page uint, pageSize uint) *repository.PageModel {
	return repository.Repository.Device.GetUserDevices(l.GetDefaultDb(), userId, page, pageSize)
}

func (l Device) GetUserDeviceById(userId model.UID, id uint) (*model.Device, error) {
	device := repository.Repository.Device.GetUserDeviceById(l.GetDefaultDb(), userId, id)
	if device == nil {
		return nil, exception.NewResponseError("设备不存在")
	}

	return device, nil
}

func (l Device) UpdateDevice(ctx context.Context, device *model.Device) error {
	if !repository.Repository.Device.UpdateDevice(l.GetDefaultDb(), device) {
		return exception.NewResponseError("更新设备失败")
	}

	return nil
}

func (l Device) OnOnlineStatucChange(device *model.Device, lastIp string, isOnline bool) error {
	if isOnline {
		device.OnlineStatus = model.DeviceOnline
		device.LastIp = lastIp
		device.LatestVisit = time.Now().Format(model.TimeFormat)
	} else {
		device.OnlineStatus = model.DeviceOffline
	}
	err := Logic.Device.UpdateDevice(context.Background(), device)
	app.GApp.GetEvent().Publish(reflect.TypeOf(event.DeviceStatusChange{}).Name(), event.NewDeviceStatusChangeEvent(device))

	return err
}

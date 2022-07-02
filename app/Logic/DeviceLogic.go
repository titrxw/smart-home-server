package logic

import (
	"context"
	"errors"

	"github.com/titrxw/smart-home-server/app/Adapter/Interface"
	"github.com/titrxw/smart-home-server/config"

	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
	"gorm.io/gorm"
)

type DeviceLogic struct {
	SupportDeviceMap     map[string]config.Device
	SupportDeviceAdapter map[string]Interface.DeviceAdapterInterface

	LogicAbstract
}

func (deviceLogic *DeviceLogic) RegisterDevice(deviceType string, device config.Device) {
	deviceLogic.SupportDeviceMap[deviceType] = device
}

func (deviceLogic DeviceLogic) GetDeviceSupportMap() map[string]config.Device {
	return deviceLogic.SupportDeviceMap
}

func (deviceLogic *DeviceLogic) RegisterDeviceAdapter(deviceType string, adapterInterface Interface.DeviceAdapterInterface) {
	deviceLogic.SupportDeviceAdapter[deviceType] = adapterInterface
}

func (deviceLogic DeviceLogic) GetDeviceAdapter(deviceType string) Interface.DeviceAdapterInterface {
	return deviceLogic.SupportDeviceAdapter[deviceType]
}

func (deviceLogic DeviceLogic) GetDeviceByDeviceId(deviceId string) *model.Device {
	app := repository.AppRepository{}.GetByAppId(deviceLogic.GetDefaultDb(), deviceId)
	if app == nil {
		return nil
	}
	return repository.DeviceRepository{}.GetDeviceByApp(deviceLogic.GetDefaultDb(), app)
}

func (deviceLogic DeviceLogic) GetDeviceById(deviceId uint) *model.Device {
	return repository.DeviceRepository{}.GetDeviceById(deviceLogic.GetDefaultDb(), deviceId)
}

func (deviceLogic DeviceLogic) CreateUserDevice(ctx context.Context, userId model.UID, deviceName string, deviceType string) (*model.Device, error) {
	device := &model.Device{
		Name: deviceName,
		Type: deviceType,
	}

	err := deviceLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		app := repository.Repository.AppRepository.CreateDeviceApp(tx)
		if app == nil {
			return errors.New("创建app失败")
		}

		device = repository.Repository.DeviceRepository.AddUserDeviceByApp(tx, userId, app, device)
		if device == nil {
			return errors.New("创建设备失败")
		}

		return Logic.EmqxLogic.AddEmqxClient(ctx, device)
	})
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (deviceLogic DeviceLogic) GetUserDevices(userId model.UID, page uint, pageSize uint) *repository.PageModel {
	return repository.DeviceRepository{}.GetUserDevices(deviceLogic.GetDefaultDb(), userId, page, pageSize)
}

func (deviceLogic DeviceLogic) GetUserDeviceById(userId model.UID, id uint) (*model.Device, error) {
	device := repository.Repository.DeviceRepository.GetUserDeviceById(deviceLogic.GetDefaultDb(), userId, id)
	if device == nil {
		return nil, errors.New("设备不存在")
	}

	return device, nil
}

func (deviceLogic DeviceLogic) UpdateDevice(ctx context.Context, device *model.Device) error {
	return deviceLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		if !repository.Repository.DeviceRepository.UpdateDevice(tx, device) {
			return errors.New("更新设备失败")
		}

		if device.IsDelete() {
			return Logic.EmqxLogic.DeleteEmqxClient(ctx, device)
		}
		return nil
	})
}

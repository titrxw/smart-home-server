package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	exception "github.com/titrxw/smart-home-server/app/Exception"

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

func (deviceLogic *DeviceLogic) RegisterDeviceAdapter(adapterInterface Interface.DeviceAdapterInterface) {
	if deviceLogic.SupportDeviceMap == nil {
		deviceLogic.SupportDeviceMap = make(map[string]config.Device)
	}
	if deviceLogic.SupportDeviceAdapter == nil {
		deviceLogic.SupportDeviceAdapter = make(map[string]Interface.DeviceAdapterInterface)
	}

	deviceConfig := adapterInterface.GetDeviceConfig()
	deviceLogic.SupportDeviceAdapter[deviceConfig.Type] = adapterInterface
	deviceLogic.SupportDeviceMap[deviceConfig.Type] = deviceConfig
}

func (deviceLogic DeviceLogic) GetDeviceSupportMap() map[string]config.Device {
	return deviceLogic.SupportDeviceMap
}

func (deviceLogic DeviceLogic) GetDeviceSupportOperateMap(device *model.Device) []string {
	_, exists := deviceLogic.GetDeviceSupportMap()[device.Type]
	if !exists {
		return nil
	}

	return deviceLogic.GetDeviceSupportMap()[device.Type].SupportOperate
}

func (deviceLogic DeviceLogic) IsSupportOperate(device *model.Device, operate model.DeviceOperateType) bool {
	supportOperateMap := deviceLogic.GetDeviceSupportOperateMap(device)
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

func (deviceLogic DeviceLogic) GetDeviceSupportReportMap(device *model.Device) []string {
	_, exists := deviceLogic.GetDeviceSupportMap()[device.Type]
	if !exists {
		return nil
	}

	return deviceLogic.GetDeviceSupportMap()[device.Type].SupportReport
}

func (deviceLogic DeviceLogic) IsSupportReport(device *model.Device, operate model.DeviceReportType) bool {
	supportReportMap := deviceLogic.GetDeviceSupportReportMap(device)
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

func (deviceLogic DeviceLogic) GetDeviceAdapter(deviceType string) Interface.DeviceAdapterInterface {
	return deviceLogic.SupportDeviceAdapter[deviceType]
}

func (deviceLogic DeviceLogic) GetDeviceByDeviceId(deviceId string) *model.Device {
	app := Logic.AppLogic.GetAppByAppId(deviceId)
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
			return exception.NewLogicError("创建app失败")
		}

		device = repository.Repository.DeviceRepository.AddUserDeviceByApp(tx, userId, app, device)
		if device == nil {
			return exception.NewLogicError("创建设备失败")
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
		return nil, exception.NewLogicError("设备不存在")
	}

	return device, nil
}

func (deviceLogic DeviceLogic) UpdateDevice(ctx context.Context, device *model.Device) error {
	return deviceLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		if !repository.Repository.DeviceRepository.UpdateDevice(tx, device) {
			return exception.NewLogicError("更新设备失败")
		}

		if device.IsDelete() {
			return Logic.EmqxLogic.DeleteEmqxClient(ctx, device)
		}
		return nil
	})
}

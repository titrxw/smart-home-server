package logic

import (
	"context"
	"reflect"
	"time"

	global "github.com/titrxw/go-framework/src/Global"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	event "github.com/titrxw/smart-home-server/app/Event"
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
	deviceLogic.SupportDeviceAdapter[deviceConfig.TypeName] = adapterInterface
	deviceLogic.SupportDeviceMap[deviceConfig.TypeName] = deviceConfig
}

func (deviceLogic DeviceLogic) GetDeviceSupportMap() map[string]config.Device {
	return deviceLogic.SupportDeviceMap
}

func (deviceLogic DeviceLogic) GetDeviceSupportOperateMap(device *model.Device) []string {
	return deviceLogic.GetDeviceSupportMap()[device.TypeName].SupportOperate
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
	return deviceLogic.GetDeviceSupportMap()[device.TypeName].SupportReport
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

func (deviceLogic DeviceLogic) IsNeedGateway(device *model.Device) bool {
	return deviceLogic.GetDeviceSupportMap()[device.TypeName].NeedGateway
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
		Name:     deviceName,
		TypeName: deviceType,
		Type:     deviceLogic.GetDeviceSupportMap()[deviceType].Type,
	}

	err := deviceLogic.GetDefaultDb().Transaction(func(tx *gorm.DB) error {
		app := repository.Repository.AppRepository.CreateDeviceApp(tx, device.Type)
		if app == nil {
			return exception.NewLogicError("创建app失败")
		}

		device = repository.Repository.DeviceRepository.AddUserDeviceByApp(tx, userId, app, device)
		if device == nil {
			return exception.NewLogicError("创建设备失败")
		}

		if !deviceLogic.GetDeviceSupportMap()[deviceType].NeedGateway {
			adapter := deviceLogic.GetDeviceAdapter(deviceType)
			return Logic.EmqxLogic.AddEmqxClient(ctx, device, map[string][]string{
				"pub": []string{
					adapter.GetReportTopic(app.AppId),
					adapter.GetAvailabilityTopic(app.AppId),
				},
				"sub": []string{
					adapter.GetCtrlTopic(app.AppId, "+"),
				},
			})
		}

		return nil
	})

	return device, err
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
	if !repository.Repository.DeviceRepository.UpdateDevice(deviceLogic.GetDefaultDb(), device) {
		return exception.NewLogicError("更新设备失败")
	}

	return nil
}

func (deviceLogic DeviceLogic) OnOnlineStatucChange(device *model.Device, lastIp string, isOnline bool) error {
	if isOnline {
		device.OnlineStatus = model.DEVICE_ONLINE
		device.LastIp = lastIp
		device.LatestVisit = time.Now().Format(model.TimeFormat)
	} else {
		device.OnlineStatus = model.DEVICE_OFFLINE
	}
	err := Logic.DeviceLogic.UpdateDevice(context.Background(), device)
	global.FApp.Event.Publish(reflect.TypeOf(event.DeviceStatusChangeEvent{}).Name(), event.NewDeviceStatusChangeEvent(device))

	return err
}

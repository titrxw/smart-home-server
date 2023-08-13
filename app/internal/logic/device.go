package logic

import (
	"context"
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	"github.com/titrxw/smart-home-server/app/internal/event"
	"github.com/titrxw/smart-home-server/app/internal/http"
	"github.com/titrxw/smart-home-server/app/internal/model"
	"github.com/titrxw/smart-home-server/app/internal/repository"
	"github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
	pkgmodel "github.com/titrxw/smart-home-server/app/pkg/model"
	pkgrepository "github.com/titrxw/smart-home-server/app/pkg/repository"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Device struct {
	logic.Abstract
}

func (l Device) IsSupportOperate(device *model.Device, operate model.DeviceOperateType) bool {
	supportOperateMap := manager.GetDeviceConfigByDeviceType(device.TypeName).SupportOperate
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
	supportReportMap := manager.GetDeviceConfigByDeviceType(device.TypeName).SupportReport
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
	deviceApp := Logic.App.GetDeviceAppByAppId(deviceId)
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
		deviceApp := repository.Repository.App.CreateApp(tx, model.DeviceAppType)
		if deviceApp == nil {
			return exception.NewResponseError("创建app失败")
		}

		device = repository.Repository.Device.AddUserDeviceByApp(tx, userId, deviceApp, device)
		if device == nil {
			return exception.NewResponseError("创建设备失败")
		}

		if !manager.GetDeviceSupportMap()[deviceType].NeedGateway {
			deviceConfig := manager.GetDeviceConfigByDeviceType(deviceType)
			subTopic := strings.Replace(deviceConfig.CtrlTopic, "{appid}", deviceApp.AppId, 1)
			subTopic = strings.Replace(subTopic, "{component_appid}", "+", 1)

			return Logic.Emqx.AddEmqxClient(ctx, device, map[string][]string{
				"pub": {
					strings.Replace(deviceConfig.ReportTopic, "{appid}", deviceApp.AppId, 1),
					strings.Replace(deviceConfig.AvailabilityTopic, "{appid}", deviceApp.AppId, 1),
				},
				"sub": {
					subTopic,
				},
			})
		}

		return nil
	})

	return device, err
}

func (l Device) GetUserDevices(userId model.UID, page uint, pageSize uint) *pkgrepository.PageModel {
	return repository.Repository.Device.GetUserDevices(l.GetDefaultDb(), userId, page, pageSize)
}

func (l Device) GetUserDeviceById(userId model.UID, id uint) (*model.Device, error) {
	device := repository.Repository.Device.GetUserDeviceById(l.GetDefaultDb(), userId, id)
	if device == nil {
		return nil, exception.NewResponseError("设备不存在")
	}

	return device, nil
}

func (l Device) GetUserDeviceByDeviceId(userId model.UID, deviceId string) (*model.Device, error) {
	deviceApp := Logic.App.GetDeviceAppByAppId(deviceId)
	if deviceApp == nil {
		return nil, exception.NewResponseError("设备不存在")
	}

	device := repository.Repository.Device.GetUserDeviceByApp(l.GetDefaultDb(), userId, deviceApp)
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

func (l Device) OnOnlineStatusChange(device *model.Device, lastIp string, isOnline bool) error {
	if isOnline {
		device.OnlineStatus = model.DeviceOnline
		device.LastIp = lastIp
		device.LatestVisit = time.Now().Format(pkgmodel.TimeFormat)
	} else {
		device.OnlineStatus = model.DeviceOffline
	}
	err := Logic.Device.UpdateDevice(context.Background(), device)
	app.GApp.GetEvent().Publish(reflect.TypeOf(event.DeviceStatusChange{}).Name(), event.NewDeviceStatusChangeEvent(device))

	return err
}

// 推动上报信息到对应的设备
func (l Device) PushMsgToDeviceServer(ctx context.Context, gatewayDevice *model.Device, device *model.Device, iotMessage *device.OperateMessage) {
	config := manager.GetDeviceConfigByDeviceType(device.TypeName)
	data, exists := config.Setting["report_http_url"]
	if !exists {
		return
	}
	reportPushHttpUrl, ok := data.(string)
	if !ok {
		return
	}

	appid := ""
	data, exists = config.Setting["report_appid"]
	if exists {
		appid, ok = data.(string)
		if !ok {
			appid = ""
		}
	}
	appSecret := ""
	data, exists = config.Setting["report_appsecret"]
	if exists {
		appSecret, ok = data.(string)
		if !ok {
			appSecret = ""
		}
	}

	pushData := map[string]interface{}{
		"gateway_appid": gatewayDevice.App.AppId,
		"device_appid":  device.App.AppId,
		"message":       iotMessage,
	}
	_, err := http.PostWithAppSignByWWWForm(ctx, appid, appSecret, reportPushHttpUrl, pushData, nil)
	if err != nil {
		helper.ErrLog(err)
		return
	}
}

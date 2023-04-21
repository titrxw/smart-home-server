package model

import deviceInterface "github.com/titrxw/smart-home-server/app/devices/interface"

type DeviceType string
type DeviceOperateType string
type DeviceReportType string

const (
	DeviceEnable  uint8 = 1
	DeviceDisable uint8 = 2
	DeviceDelete  uint8 = 3
)

const (
	DeviceOnline  = 1
	DeviceOffline = 2
)

type Device struct {
	Model

	Name            string    `json:"name" gorm:"type:varchar(64);not null"`
	UserId          UID       `json:"user_id" gorm:"not null"`
	AppId           uint      `json:"-" gorm:"not null"`
	Type            uint8     `json:"type" gorm:"not null;default:1"`
	TypeName        string    `json:"type_name" gorm:"type:varchar(64);not null"`
	GatewayDeviceId uint      `json:"-" gorm:"not null;default:0"`
	LatestVisit     string    `json:"latest_visit" gorm:"type:varchar(12);not null;default:''"`
	OnlineStatus    uint8     `json:"online_status" gorm:"not null;default:0"`
	LastIp          string    `json:"last_ip" gorm:"type:varchar(20);not null;default:''"`
	DeviceStatus    uint8     `json:"device_status" gorm:"not null;default:1"`
	DeviceCurStatus string    `json:"device_cur_status" gorm:"type:varchar(500);not null;default:''"`
	CreatedAt       LocalTime `json:"created_at"`

	App               *App               `json:"-"`
	DeviceOperateLogs []DeviceOperateLog `json:"-"`
	GatewayDevice     *Device            `json:"gateway_device"`
}

func (device *Device) Enable() {
	device.DeviceStatus = DeviceEnable
}

func (device *Device) Disable() {
	device.DeviceStatus = DeviceDisable
}

func (device *Device) IsDisable() bool {
	return device.DeviceStatus == DeviceDisable
}

func (device *Device) Delete() {
	device.DeviceStatus = DeviceDelete
}

func (device *Device) IsDelete() bool {
	return device.DeviceStatus == DeviceDelete
}

func (device *Device) IsOnline() bool {
	return device.OnlineStatus == DeviceOnline
}

func (device *Device) IsGateway() bool {
	return device.Type == deviceInterface.DeviceGatewayAppType
}

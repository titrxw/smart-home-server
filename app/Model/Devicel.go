package model

type DeviceType string
type DeviceOperateType string
type DeviceReportType string

const DEVICE_APP_TYPE uint8 = 1
const DEVICE_GATEWAY_APP_TYPE uint8 = 2

const (
	DEVICE_ENABLE  uint8 = 1
	DEVICE_DISABLE uint8 = 2
	DEVICE_DELETE  uint8 = 3
)

const (
	DEVICE_ONLINE  = 1
	DEVICE_OFFLINE = 2
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
	device.DeviceStatus = DEVICE_ENABLE
}

func (device *Device) Disable() {
	device.DeviceStatus = DEVICE_DISABLE
}

func (device *Device) IsDisable() bool {
	return device.DeviceStatus == DEVICE_DISABLE
}

func (device *Device) Delete() {
	device.DeviceStatus = DEVICE_DELETE
}

func (device *Device) IsDelete() bool {
	return device.DeviceStatus == DEVICE_DELETE
}

func (device *Device) IsOnline() bool {
	return device.OnlineStatus == DEVICE_ONLINE
}

func (device *Device) IsGateway() bool {
	return device.Type == DEVICE_GATEWAY_APP_TYPE
}

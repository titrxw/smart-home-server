package model

type DeviceType string
type DeviceOperateType string

const (
	DEVICE_ENABLE  uint8 = 1
	DEVICE_DISABLE uint8 = 2
	DEVICE_DELETE  uint8 = 3
)

type Device struct {
	Model

	Name            string    `json:"name" gorm:"type:varchar(64);not null"`
	UserId          UID       `json:"user_id" gorm:"not null"`
	AppId           uint      `json:"-" gorm:"not null"`
	Type            string    `json:"type" gorm:"type:varchar(12);not null"`
	LatestVisit     string    `json:"latest_visit" gorm:"type:varchar(12);not null;default:''"`
	OnlineStatus    uint8     `json:"online_status" gorm:"not null;default:0"`
	DeviceStatus    uint8     `json:"device_status" gorm:"not null;default:1"`
	DeviceCurStatus string    `json:"device_cur_status" gorm:"type:varchar(500);not null;default:''"`
	CreatedAt       LocalTime `json:"created_at"`

	App               *App               `json:"-"`
	DeviceOperateLogs []DeviceOperateLog `json:"-"`
}

func (this *Device) Enable() {
	this.DeviceStatus = DEVICE_ENABLE
}

func (this *Device) Disable() {
	this.DeviceStatus = DEVICE_DISABLE
}

func (this *Device) IsDisable() bool {
	return this.DeviceStatus == DEVICE_DISABLE
}

func (this *Device) Delete() {
	this.DeviceStatus = DEVICE_DELETE
}

func (this *Device) IsDelete() bool {
	return this.DeviceStatus == DEVICE_DELETE
}

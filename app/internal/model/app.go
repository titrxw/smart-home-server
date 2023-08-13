package model

import "github.com/titrxw/smart-home-server/app/pkg/model"

const DeviceAppType uint8 = 1
const OpenAppType uint8 = 2

type App struct {
	model.Model

	AppId     string `json:"appid" gorm:"type:varchar(32);not null;uniqueIndex"`
	AppSecret string `json:"-" gorm:"type:varchar(64);not null"`
	AppType   uint8  `json:"-" gorm:"not null;default:1;comment:1代表设备，2代表对外开放的app"`
}

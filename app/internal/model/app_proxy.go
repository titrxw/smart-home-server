package model

import "github.com/titrxw/smart-home-server/app/pkg/model"

type AppProxy struct {
	model.Model

	AppId          string `json:"appid" gorm:"not null"`
	ComponentAppId string `json:"component_appid" gorm:"not null"`
}

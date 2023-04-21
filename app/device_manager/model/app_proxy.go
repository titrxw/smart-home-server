package model

type AppProxy struct {
	Model

	AppId          string `json:"appid" gorm:"not null"`
	ComponentAppId string `json:"component_appid" gorm:"not null"`
}

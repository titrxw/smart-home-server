package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/titrxw/smart-home-server/app/pkg/model"
)

type FaceUrls []string

func (payload FaceUrls) Value() (driver.Value, error) {
	if payload == nil {
		return nil, nil
	}

	return json.Marshal(payload)
}

func (payload *FaceUrls) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}
	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, &payload)
}

const FaceModelStatusEnable = 1
const FaceModelStatusDisable = 0

type FaceModel struct {
	model.Model

	UserId      uint     `json:"user_id" gorm:"not null"`
	DeviceAppId string   `json:"device_appid" gorm:"column:device_appid;not null"`
	UserName    string   `json:"user_name" gorm:"type:varchar(64);not null"`
	Status      uint8    `json:"status" gorm:"type:tinyint(4);not null"`
	Urls        FaceUrls `json:"face_urls" gorm:"type:text;not null;default:''"`

	CreatedAt model.LocalTime `json:"created_at"`
}

func (m *FaceModel) IsEnable() bool {
	return m.Status == FaceModelStatusEnable
}

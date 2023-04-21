package repository

import (
	"encoding/json"
	"github.com/titrxw/smart-home-server/app/device_manager/model"

	"gorm.io/gorm"
)

type Setting struct {
	Abstract
}

func (r Setting) Set(db *gorm.DB, key string, value interface{}) bool {
	data, _ := json.Marshal(value)
	settingModel := &model.Setting{
		Key:   key,
		Value: string(data),
	}

	result := db.Create(settingModel)
	if result.RowsAffected == 1 {
		return true
	}

	return false
}

func (r Setting) Get(db *gorm.DB, key string, defaultValue interface{}) interface{} {
	settingModel := new(model.Setting)
	result := db.Where("key = ?", key).First(settingModel)
	if result.RowsAffected == 1 {
		var data interface{}
		err := json.Unmarshal([]byte(settingModel.Value), &data)
		if err == nil {
			return data
		}
	}

	return defaultValue
}

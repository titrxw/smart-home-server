package repository

import (
	"encoding/json"

	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
)

type SettingRepository struct {
	RepositoryAbstract
}

func (settingRepository SettingRepository) Set(db *gorm.DB, key string, value interface{}) bool {
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

func (settingRepository SettingRepository) Get(db *gorm.DB, key string, defaultValue interface{}) interface{} {
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

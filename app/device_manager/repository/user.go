package repository

import (
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Abstract
}

func (r User) CreateUser(db *gorm.DB, user *model.User) bool {
	user.Salt = helper.RandomStr(12)
	user.Password = user.MakeHashPassword(user.Password, user.Salt)
	user.RegisterAt = model.LocalTime(time.Now())
	user.CreatedAt = model.LocalTime(time.Now())

	result := db.Create(user)

	return result.RowsAffected == 1
}

func (r User) GetById(db *gorm.DB, id model.UID) *model.User {
	user := new(model.User)
	result := db.Where("id = ?", id).First(user)
	if result.RowsAffected == 1 {
		return user
	}

	return nil
}

func (r User) GetByUserName(db *gorm.DB, userName string) *model.User {
	user := new(model.User)
	result := db.Where("user_name = ?", userName).First(user)
	if result.RowsAffected == 1 {
		return user
	}

	return nil
}

func (r User) GetByEmail(db *gorm.DB, email string) *model.User {
	user := new(model.User)
	result := db.Where("email = ?", email).First(user)
	if result.RowsAffected == 1 {
		return user
	}

	return nil
}

func (r User) GetByMobile(db *gorm.DB, mobile string) *model.User {
	user := new(model.User)
	result := db.Where("mobile = ?", mobile).First(user)
	if result.RowsAffected == 1 {
		return user
	}

	return nil
}

func (r User) UpdateUser(db *gorm.DB, user *model.User) bool {
	result := db.Save(user)

	return result.Error == nil
}

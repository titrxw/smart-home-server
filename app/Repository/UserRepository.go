package repository

import (
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	RepositoryAbstract
}

func (this UserRepository) CreateUser(db *gorm.DB, user *model.User) bool {
	user.Salt = helper.RandomStr(12)
	user.Password = user.MakeHashPassword(user.Password, user.Salt)
	user.RegisterAt = model.LocalTime(time.Now())
	user.CreatedAt = model.LocalTime(time.Now())

	result := db.Create(user)

	return result.RowsAffected == 1
}

func (this UserRepository) GetById(db *gorm.DB, id model.UID) *model.User {
	user := new(model.User)
	result := db.Where("id = ?", id).First(user)
	if result.RowsAffected == 1 {
		return user
	}

	return nil
}

func (this UserRepository) GetByUserName(db *gorm.DB, userName string) *model.User {
	user := new(model.User)
	result := db.Where("user_name = ?", userName).First(user)
	if result.RowsAffected == 1 {
		return user
	}

	return nil
}

func (this UserRepository) GetByMobile(db *gorm.DB, mobile string) *model.User {
	user := new(model.User)
	result := db.Where("mobile = ?", mobile).First(user)
	if result.RowsAffected == 1 {
		return user
	}

	return nil
}

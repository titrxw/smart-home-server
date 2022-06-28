package model

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type UID uint

const PWD_AUTH_CODE = "*mDG6HNY09*VY3AL"
const (
	USER_ENABLE  uint8 = 1
	USER_DISABLE uint8 = 0
)

type User struct {
	Model

	UserName   string    `json:"user_name" gorm:"type:varchar(32);not null"`
	Mobile     string    `json:"mobile" gorm:"type:varchar(11);not null;uniqueIndex"`
	Password   string    `json:"-" gorm:"type:varchar(64);not null"`
	Salt       string    `json:"-" gorm:"type:varchar(12);not null"`
	Status     uint8     `json:"status" gorm:"not null;default:1"`
	CreatedAt  LocalTime `json:"created_at"`
	RegisterAt LocalTime `json:"register_at"`

	Devices []Device `json:"-"`
}

func (this User) MakeHashPassword(password string, salt string) string {
	password = password + "-" + salt + "-" + PWD_AUTH_CODE
	hash := sha1.New()
	_, _ = io.WriteString(hash, password)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (this *User) IsDisable() bool {
	return this.Status == USER_DISABLE
}

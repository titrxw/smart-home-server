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

	UserName    string    `json:"user_name" gorm:"type:varchar(32);not null"`
	Mobile      string    `json:"mobile" gorm:"type:varchar(11);not null;uniqueIndex"`
	Password    string    `json:"-" gorm:"type:varchar(64);not null"`
	Salt        string    `json:"-" gorm:"type:varchar(12);not null"`
	Status      uint8     `json:"status" gorm:"not null;default:1"`
	LastIp      string    `json:"last_ip" gorm:"type:varchar(20);not null;default:''"`
	LatestVisit string    `json:"latest_visit" gorm:"type:varchar(12);not null;default:''"`
	CreatedAt   LocalTime `json:"created_at"`
	RegisterAt  LocalTime `json:"register_at"`

	Devices []Device `json:"-"`
}

func (user User) MakeHashPassword(password string, salt string) string {
	password = password + "-" + salt + "-" + PWD_AUTH_CODE
	hash := sha1.New()
	_, _ = io.WriteString(hash, password)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (user *User) IsDisable() bool {
	return user.Status == USER_DISABLE
}

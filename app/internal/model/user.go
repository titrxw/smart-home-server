package model

import (
	"crypto/sha1"
	"fmt"
	"github.com/titrxw/smart-home-server/app/pkg/model"
	"io"
)

type UID uint

const PwdAuthCode = "*mDG6HNY09*VY3AL"
const (
	UserDisable uint8 = 0
)

type User struct {
	model.Model

	UserName    string          `json:"user_name" gorm:"type:varchar(32);not null"`
	Email       string          `json:"email" gorm:"type:varchar(32);not null;uniqueIndex"`
	Mobile      string          `json:"mobile" gorm:"type:varchar(11);not null;default:'';uniqueIndex"`
	Password    string          `json:"-" gorm:"type:varchar(64);not null"`
	Salt        string          `json:"-" gorm:"type:varchar(12);not null"`
	Status      uint8           `json:"status" gorm:"not null;default:1"`
	LastIp      string          `json:"last_ip" gorm:"type:varchar(20);not null;default:''"`
	LatestVisit string          `json:"latest_visit" gorm:"type:varchar(12);not null;default:''"`
	CreatedAt   model.LocalTime `json:"created_at"`
	RegisterAt  model.LocalTime `json:"register_at"`

	Devices []Device `json:"-"`
}

func (user *User) MakeHashPassword(password string, salt string) string {
	password = password + "-" + salt + "-" + PwdAuthCode
	hash := sha1.New()
	_, _ = io.WriteString(hash, password)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (user *User) IsDisable() bool {
	return user.Status == UserDisable
}

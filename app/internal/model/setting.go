package model

import "github.com/titrxw/smart-home-server/app/pkg/model"

type Setting struct {
	model.Model

	Key   string `gorm:"type:varchar(120);not null;uniqueIndex"`
	Value string `gorm:"type:text;not null"`
}

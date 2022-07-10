package model

type Setting struct {
	Model

	Key   string `gorm:"type:varchar(120);not null;uniqueIndex"`
	Value string `gorm:"type:text;not null"`
}

package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID   int
	Message  string `gorm:"column:message;varchar(256);not null"`
	Pictures string `gorm:"column:picture_path;varchar(256);"`
}

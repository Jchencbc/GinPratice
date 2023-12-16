package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"column:name;varchar(20);not null;unique"`
	Telephone   string `gorm:"column:telephone;varchar(20);not null;unique"`
	Password    string `gorm:"column:password;size:255;not null"`
	IsSuperuser int8   `gorm:"column:is_superuser;not null;default 0"`
}

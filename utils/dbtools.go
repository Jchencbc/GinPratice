package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:zz.@tcp(zz:3306)/zz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	// db.AutoMigrate(&models.User{}) //迁移model层
	return db
}

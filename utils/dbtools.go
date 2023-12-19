package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:48mysqljc.@tcp(1.94.40.252:3306)/research_platform_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	// db.AutoMigrate(&models.User{}) //迁移model层
	return db
}

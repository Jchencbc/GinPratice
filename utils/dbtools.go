package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// driverName := "mysql"
	// host := "139.196.100.132"
	// port := "3306"
	// database := "research_platform_go"
	// username := "qinghua"
	// password := "qinghua"
	// charset := "utf8mb4"
	// args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
	// 	username,
	// 	password,
	// 	host,
	// 	port,
	// 	database,
	// 	charset)
	dsn := "root:jB5OgGyIIBO@tcp(1.94.40.252:3306)/research_platform_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	// db.AutoMigrate(&models.User{})
	return db
}

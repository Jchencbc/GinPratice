package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	cfg, _ := ini.Load("../config/app.ini")
	host := cfg.Section("Mysql").Key("host").String()
	port := cfg.Section("Mysql").Key("port").String()
	database := cfg.Section("Mysql").Key("database").String()
	username := cfg.Section("Mysql").Key("username").String()
	password := cfg.Section("Mysql").Key("password").String()
	charset := cfg.Section("Mysql").Key("charset").String()
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	// db.AutoMigrate(&models.User{}) //迁移model层
	return DB
}

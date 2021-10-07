package config

import (
	"echo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/bunn_petshop?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Pet{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Category{})
	DB.AutoMigrate(&model.Article{})
}

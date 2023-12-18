package config

import (
	"capstone/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigrate()
	return DB
}

func initMigrate() {
	err := DB.AutoMigrate(&model.User{}, &model.Plant{}, &model.Bookmark{}, &model.Like{}, &model.ScanPlant{})
	if err != nil {
		panic("failed to migrate database")
	}
}

package service

import (
	"server/mdl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

const (
	dsn = ""
)

func init() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&model.User{}, &model.Work{}, &model.Drink{}); err != nil {
		panic("failed to migrate table")
	}
}

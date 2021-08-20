package service

import (
	. "server/mdl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"errors"
)

var Db *gorm.DB

const (
	dsn = "add:password@tcp(127.0.0.1:3306)/hogehoge?charset=utf8mb4&parseTime=True"
)

func init() {
	err := errors.New("")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db = Db.Debug()
	if err != nil {
		panic("failed to connect database")
	}
	if err := Db.AutoMigrate(&User{}, &Work{}, &Drink{}); err != nil {
		panic("failed to migrate table")
	}
}

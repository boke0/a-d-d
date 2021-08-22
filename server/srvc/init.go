package service

import (
	. "server/mdl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"errors"
	"os"
)

var Db *gorm.DB

var (
	dsn = os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/"+os.Getenv("MYSQL_DATABASE")+"?charset=utf8mb4&parseTime=True"
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

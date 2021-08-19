package model

import (
	"time"
	"gorm.io/gorm"
)

type Work struct {
	gorm.Model
	Title string
	Description string
	UserId uint
	User User
	StartTime time.Time `gorm:"type:datetime(6)"` 
	EndTime time.Time `gorm:"type:datetime(6)"` 
	Drinks []Drink
}

type CreateWorkParam struct {
	Title string
	Description string
	StartTime time.Time
	Drinks []Drink
}

type UpdateWorkParam struct {
	Title string
	Description string
	EndTime time.Time
}

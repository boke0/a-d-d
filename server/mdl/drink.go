package model

import (
	"gorm.io/gorm"
)

type Drink struct {
	gorm.Model
	Name string
	Description string
	WorkId uint
	Work Work
	Image string
	Alcohol float32
	Amount float32
}

type CreateDrinkParam struct {
	Name string
	Description string
	Image string
	Alcohol float32
	Amount float32
}

type UpdateDrinkParam struct {
	Name string
	Description string
	Image string
	Alcohol float32
	Amount float32
}

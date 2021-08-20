package model

import (
	"gorm.io/gorm"
)

type Drink struct {
	gorm.Model
	Name string
	WorkId uint
	Work Work
	Alcohol float32
	Amount float32
}

type CreateDrinkParam struct {
	Name string
	Alcohol float32
	Amount float32
}

type UpdateDrinkParam struct {
	Name string
	Alcohol float32
	Amount float32
}

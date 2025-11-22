package domain

import "gorm.io/gorm"

type OrderFood struct {
	gorm.Model
	OrderID  uint `gorm:"primaryKey"`
	FoodID   uint `gorm:"primaryKey"`
	Quantity int
	Notes    string
}

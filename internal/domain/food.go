package domain

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name  string
	Price float64
}

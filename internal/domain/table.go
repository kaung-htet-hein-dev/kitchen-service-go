package domain

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Number int
	Orders []Order `gorm:"foreignKey:TableID"`
}

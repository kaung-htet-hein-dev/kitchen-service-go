package main

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/domain"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/infrastructure/db"

	"gorm.io/gorm"
)

func SeedFoodsAndTables(db *gorm.DB) error {
	foods := []domain.Food{
		{Name: "Pizza", Price: 12.99},
		{Name: "Burger", Price: 8.49},
		{Name: "Pasta", Price: 10.25},
		{Name: "Salad", Price: 6.75},
		{Name: "Sushi", Price: 15.00},
	}
	tables := []domain.Table{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
	}
	if err := db.Create(&foods).Error; err != nil {
		return err
	}
	if err := db.Create(&tables).Error; err != nil {
		return err
	}
	return nil
}

func main() {
	db := db.ConnectDB()

	if err := SeedFoodsAndTables(db); err != nil {
		panic(err)
	}
}

package usecase

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/infrastructure/repository"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"
)

type FoodUsecase struct {
	foodRepo *repository.FoodHandler
}

func NewFoodUsecase(foodRepo *repository.FoodHandler) *FoodUsecase {
	return &FoodUsecase{foodRepo: foodRepo}
}

func (u *FoodUsecase) GetFoods() ([]response.FoodResponse, error) {
	foods, err := u.foodRepo.GetFoods()
	if err != nil {
		return nil, err
	}
	return foods, nil
}

func (u *FoodUsecase) CreateFood() {
	// Business logic to create a food item
}
func (u *FoodUsecase) UpdateFood() {
	// Business logic to update a food item
}

func (u *FoodUsecase) DeleteFood() {
	// Business logic to delete a food item
}

package usecase

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/infrastructure/repository"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"
)

type FoodUsecase struct {
	foodRepo *repository.FoodRepository
}

func NewFoodUsecase(foodRepo *repository.FoodRepository) *FoodUsecase {
	return &FoodUsecase{foodRepo: foodRepo}
}

func (u *FoodUsecase) GetFoods() ([]response.FoodResponse, error) {
	foods, err := u.foodRepo.GetFoods()
	if err != nil {
		return nil, err
	}
	return foods, nil
}

func (u *FoodUsecase) GetFood(id string) (*response.FoodResponse, error) {
	return u.foodRepo.GetFood(id)
}

func (u *FoodUsecase) CreateFood(req *request.CreateFoodRequest) error {
	return u.foodRepo.CreateFood(req)
}

func (u *FoodUsecase) UpdateFood(id string, req *request.CreateFoodRequest) error {
	return u.foodRepo.UpdateFood(id, req)
}

func (u *FoodUsecase) DeleteFood(id string) error {
	return u.foodRepo.DeleteFood(id)
}

package repository

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/domain"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"

	"gorm.io/gorm"
)

type FoodHandler struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodHandler {
	return &FoodHandler{db: db}
}

func (h *FoodHandler) GetFood() error {
	return nil
}

func (h *FoodHandler) GetFoods() ([]response.FoodResponse, error) {
	var foods []response.FoodResponse
	err := h.db.Model(&domain.Food{}).Select("id, name, price").Scan(&foods).Error

	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (h *FoodHandler) CreateFood() error {
	return nil
}

func (h *FoodHandler) UpdateFood() error {
	return nil
}

func (h *FoodHandler) DeleteFood() error {
	return nil
}

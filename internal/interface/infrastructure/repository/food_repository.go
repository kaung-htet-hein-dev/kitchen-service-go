package repository

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/domain"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"
	"strconv"

	"gorm.io/gorm"
)

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{db: db}
}

func (r *FoodRepository) GetFood(id string) (*response.FoodResponse, error) {
	var food domain.Food
	if err := r.db.First(&food, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &response.FoodResponse{
		ID:    food.ID,
		Name:  food.Name,
		Price: food.Price,
	}, nil
}

func (r *FoodRepository) GetFoods() ([]response.FoodResponse, error) {
	var foods []response.FoodResponse
	err := r.db.Model(&domain.Food{}).Select("id, name, price").Scan(&foods).Error

	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (r *FoodRepository) CreateFood(req *request.CreateFoodRequest) error {
	food := domain.Food{
		Name:  req.Name,
		Price: req.Price,
	}
	return r.db.Create(&food).Error
}

func (r *FoodRepository) UpdateFood(id string, req *request.CreateFoodRequest) error {
	var food domain.Food
	if err := r.db.First(&food, id).Error; err != nil {
		return err
	}
	food.Name = req.Name
	food.Price = req.Price
	return r.db.Save(&food).Error
}

func (h *FoodRepository) DeleteFood(id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return h.db.Delete(&domain.Food{}, idUint).Error
}

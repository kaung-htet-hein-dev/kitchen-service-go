package handler

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/usecase"

	"github.com/labstack/echo/v4"
)

type FoodHandler struct {
	foodUsecase *usecase.FoodUsecase
}

func NewFoodHandler(foodUsecase *usecase.FoodUsecase) *FoodHandler {
	return &FoodHandler{foodUsecase: foodUsecase}
}

func (h *FoodHandler) GetFoods(c echo.Context) error {
	data, err := h.foodUsecase.GetFoods()

	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to get foods"})
	}

	return c.JSON(200, echo.Map{"data": data})
}

func (h *FoodHandler) CreateFood(c echo.Context, req *request.CreateFoodRequest) error {
	return nil
}

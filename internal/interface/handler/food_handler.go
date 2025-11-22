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
	err := h.foodUsecase.CreateFood(req)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to create food"})
	}
	return c.JSON(201, echo.Map{"message": "Food created"})
}

func (h *FoodHandler) GetFood(c echo.Context) error {
	id := c.Param("id")

	food, err := h.foodUsecase.GetFood(id)
	if err != nil {
		return c.JSON(404, echo.Map{"error": "Food not found"})
	}
	return c.JSON(200, echo.Map{"data": food})
}

func (h *FoodHandler) UpdateFood(c echo.Context, req *request.CreateFoodRequest) error {
	id := c.Param("id")
	err := h.foodUsecase.UpdateFood(id, req)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to update food"})
	}
	return c.JSON(200, echo.Map{"message": "Food updated"})
}

func (h *FoodHandler) DeleteFood(c echo.Context) error {
	id := c.Param("id")
	err := h.foodUsecase.DeleteFood(id)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to delete food"})
	}
	return c.JSON(200, echo.Map{"message": "Food deleted"})
}

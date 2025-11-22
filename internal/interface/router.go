package router

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/handler"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/infrastructure/repository"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/usecase"
	"kaung-htet-hein-dev/kitchen-service-go/pkg"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(e *echo.Echo, db *gorm.DB) {

}

func RegisterFoodRoutes(e *echo.Echo, db *gorm.DB) {
	foodGroup := e.Group("/foods")

	foodRepo := repository.NewFoodRepository(db)
	foodUsecase := usecase.NewFoodUsecase(foodRepo)
	handler := handler.NewFoodHandler(foodUsecase)

	foodGroup.GET("", handler.GetFoods)
	foodGroup.POST("", pkg.BindAndValidate(handler.CreateFood))
}

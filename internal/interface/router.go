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
	orderGroup := e.Group("/orders")

	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	handler := handler.NewOrderHandler(orderUsecase)

	orderGroup.GET("", handler.GetOrders)
	orderGroup.GET("/:id", handler.GetOrder)
	orderGroup.POST("", pkg.BindAndValidate(handler.CreateOrder))
	orderGroup.PUT("/:id", pkg.BindAndValidate(handler.UpdateOrder))
	orderGroup.DELETE("/:id", handler.DeleteOrder)
}

func RegisterTableRoutes(e *echo.Echo, db *gorm.DB) {
	tableGroup := e.Group("/tables")

	tableRepo := repository.NewTableRepository(db)
	tableUsecase := usecase.NewTableUsecase(tableRepo)
	handler := handler.NewTableHandler(tableUsecase)

	tableGroup.GET("", handler.GetTables)
	tableGroup.GET("/:id", handler.GetTable)
	tableGroup.POST("", pkg.BindAndValidate(handler.CreateTable))
	tableGroup.PUT("/:id", pkg.BindAndValidate(handler.UpdateTable))
	tableGroup.DELETE("/:id", handler.DeleteTable)
}

func RegisterFoodRoutes(e *echo.Echo, db *gorm.DB) {
	foodGroup := e.Group("/foods")

	foodRepo := repository.NewFoodRepository(db)
	foodUsecase := usecase.NewFoodUsecase(foodRepo)
	handler := handler.NewFoodHandler(foodUsecase)

	foodGroup.GET("", handler.GetFoods)
	foodGroup.GET("/:id", handler.GetFood)
	foodGroup.POST("", pkg.BindAndValidate(handler.CreateFood))
	foodGroup.PUT("/:id", pkg.BindAndValidate(handler.UpdateFood))
	foodGroup.DELETE("/:id", handler.DeleteFood)
}

package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(e *echo.Echo, db *gorm.DB) {
	categoryGroup := e.Group("/orders")

}

package server

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/config"
	router "kaung-htet-hein-dev/kitchen-service-go/internal/interface"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/infrastructure/db"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/middleware"
	"kaung-htet-hein-dev/kitchen-service-go/pkg"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	cfg := config.LoadConfig()

	middleware.RegisterBasicMiddleware(e)
	db := db.ConnectDB()

	router.RegisterOrderRoutes(e, db)

	log.Fatal(e.Start(":" + cfg.Port))
}

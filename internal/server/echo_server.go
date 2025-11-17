package server

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/config"
	"log"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	cfg := config.LoadConfig()

	log.Fatal(e.Start(":" + cfg.Port))
}

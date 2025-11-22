package handler

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/usecase"

	"github.com/labstack/echo/v4"
)

type TableHandler struct {
	tableUsecase *usecase.TableUsecase
}

func NewTableHandler(tableUsecase *usecase.TableUsecase) *TableHandler {
	return &TableHandler{tableUsecase: tableUsecase}
}

func (h *TableHandler) GetTables(c echo.Context) error {
	data, err := h.tableUsecase.GetTables()
	if err != nil {
		return c.JSON(500, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, echo.Map{"data": data})
}

func (h *TableHandler) GetTable(c echo.Context) error {
	id := c.Param("id")
	table, err := h.tableUsecase.GetTable(id)
	if err != nil {
		return c.JSON(404, echo.Map{"error": "Table not found"})
	}
	return c.JSON(200, echo.Map{"data": table})
}

func (h *TableHandler) CreateTable(c echo.Context, req *request.CreateTableRequest) error {
	err := h.tableUsecase.CreateTable(req)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to create table"})
	}
	return c.JSON(201, echo.Map{"message": "Table created"})
}

func (h *TableHandler) UpdateTable(c echo.Context, req *request.CreateTableRequest) error {
	id := c.Param("id")
	err := h.tableUsecase.UpdateTable(id, req)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to update table"})
	}
	return c.JSON(200, echo.Map{"message": "Table updated"})
}

func (h *TableHandler) DeleteTable(c echo.Context) error {
	id := c.Param("id")
	err := h.tableUsecase.DeleteTable(id)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to delete table"})
	}
	return c.JSON(200, echo.Map{"message": "Table deleted"})
}

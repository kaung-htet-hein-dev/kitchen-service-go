package handler

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/usecase"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

func (h *OrderHandler) GetOrders(c echo.Context) error {
	data, err := h.orderUsecase.GetOrders()
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to get orders"})
	}
	return c.JSON(200, echo.Map{"data": data})
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := h.orderUsecase.GetOrder(id)
	if err != nil {
		return c.JSON(404, echo.Map{"error": "Order not found"})
	}
	return c.JSON(200, echo.Map{"data": order})
}

func (h *OrderHandler) CreateOrder(c echo.Context, req *request.CreateOrderRequest) error {
	err := h.orderUsecase.CreateOrder(req)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to create order"})
	}
	return c.JSON(201, echo.Map{"message": "Order created"})
}

func (h *OrderHandler) UpdateOrder(c echo.Context, req *request.CreateOrderRequest) error {
	id := c.Param("id")
	err := h.orderUsecase.UpdateOrder(id, req)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to update order"})
	}
	return c.JSON(200, echo.Map{"message": "Order updated"})
}

func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	err := h.orderUsecase.DeleteOrder(id)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to delete order"})
	}
	return c.JSON(200, echo.Map{"message": "Order deleted"})
}

package usecase

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/infrastructure/repository"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"
)

type OrderUsecase struct {
	orderRepo *repository.OrderRepository
}

func NewOrderUsecase(orderRepo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: orderRepo}
}

func (u *OrderUsecase) GetOrders() ([]response.OrderResponse, error) { return u.orderRepo.GetOrders() }
func (u *OrderUsecase) GetOrder(id string) (*response.OrderResponse, error) {
	return u.orderRepo.GetOrder(id)
}
func (u *OrderUsecase) CreateOrder(req *request.CreateOrderRequest) error {
	return u.orderRepo.CreateOrder(req)
}
func (u *OrderUsecase) UpdateOrder(id string, req *request.CreateOrderRequest) error {
	return u.orderRepo.UpdateOrder(id, req)
}
func (u *OrderUsecase) DeleteOrder(id string) error { return u.orderRepo.DeleteOrder(id) }

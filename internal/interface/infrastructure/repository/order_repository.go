package repository

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/domain"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"
	"strconv"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrders() ([]response.OrderResponse, error) {
	var orders []domain.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	var resp []response.OrderResponse
	for _, o := range orders {
		resp = append(resp, response.OrderResponse{
			ID:      o.ID,
			Status:  o.Status,
			TableID: o.TableID,
		})
	}
	return resp, nil
}

func (r *OrderRepository) GetOrder(id string) (*response.OrderResponse, error) {
	var order domain.Order
	if err := r.db.First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &response.OrderResponse{
		ID:      order.ID,
		Status:  order.Status,
		TableID: order.TableID,
	}, nil
}

func (r *OrderRepository) CreateOrder(req *request.CreateOrderRequest) error {
	order := domain.Order{
		Status:  req.Status,
		TableID: req.TableID,
	}
	return r.db.Create(&order).Error
}

func (r *OrderRepository) UpdateOrder(id string, req *request.CreateOrderRequest) error {
	var order domain.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return err
	}
	order.Status = req.Status
	order.TableID = req.TableID
	return r.db.Save(&order).Error
}

func (r *OrderRepository) DeleteOrder(id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return r.db.Delete(&domain.Order{}, idUint).Error
}

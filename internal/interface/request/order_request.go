package request

type CreateOrderRequest struct {
	Status  string `json:"status" validate:"required"`
	TableID uint   `json:"table_id" validate:"required,gt=0"`
}

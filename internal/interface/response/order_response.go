package response

type OrderResponse struct {
	ID      uint   `json:"id"`
	Status  string `json:"status"`
	TableID uint   `json:"table_id"`
}

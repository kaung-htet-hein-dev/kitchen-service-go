package request

type CreateTableRequest struct {
	Number int `json:"number" validate:"required,gt=0"`
}

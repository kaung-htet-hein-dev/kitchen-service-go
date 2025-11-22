package repository

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/domain"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"
	"strconv"

	"gorm.io/gorm"
)

type TableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) *TableRepository {
	return &TableRepository{db: db}
}

func (r *TableRepository) GetTables() ([]response.TableResponse, error) {
	var tables []domain.Table
	if err := r.db.Find(&tables).Error; err != nil {
		return nil, err
	}
	var resp []response.TableResponse
	for _, t := range tables {
		resp = append(resp, response.TableResponse{
			ID:     t.ID,
			Number: t.Number,
		})
	}
	return resp, nil
}

func (r *TableRepository) GetTable(id string) (*response.TableResponse, error) {
	var table domain.Table
	if err := r.db.First(&table, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &response.TableResponse{
		ID:     table.ID,
		Number: table.Number,
	}, nil
}

func (r *TableRepository) CreateTable(req *request.CreateTableRequest) error {
	table := domain.Table{
		Number: req.Number,
	}
	return r.db.Create(&table).Error
}

func (r *TableRepository) UpdateTable(id string, req *request.CreateTableRequest) error {
	var table domain.Table
	if err := r.db.First(&table, id).Error; err != nil {
		return err
	}
	table.Number = req.Number
	return r.db.Save(&table).Error
}

func (r *TableRepository) DeleteTable(id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return r.db.Delete(&domain.Table{}, idUint).Error
}

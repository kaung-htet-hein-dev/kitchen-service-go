package usecase

import (
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/infrastructure/repository"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/request"
	"kaung-htet-hein-dev/kitchen-service-go/internal/interface/response"
)

type TableUsecase struct {
	tableRepo *repository.TableRepository
}

func NewTableUsecase(tableRepo *repository.TableRepository) *TableUsecase {
	return &TableUsecase{tableRepo: tableRepo}
}

func (u *TableUsecase) GetTables() ([]response.TableResponse, error) {
	return u.tableRepo.GetTables()
}

func (u *TableUsecase) GetTable(id string) (*response.TableResponse, error) {
	return u.tableRepo.GetTable(id)
}
func (u *TableUsecase) CreateTable(req *request.CreateTableRequest) error {
	return u.tableRepo.CreateTable(req)
}
func (u *TableUsecase) UpdateTable(id string, req *request.CreateTableRequest) error {
	return u.tableRepo.UpdateTable(id, req)
}
func (u *TableUsecase) DeleteTable(id string) error {
	return u.tableRepo.DeleteTable(id)
}

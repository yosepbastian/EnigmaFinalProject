package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type StocksUseCase interface {
	GetStocksByName(stocksName string) (models.Stocks, error)
	UpdateByName(stocks *models.Stocks) error
	GetAll() ([]models.Stocks, error)
	GetStockQtyById(stockId string) (int, error)
	UpdateQtyStockById(quantity int, id string) error
}

type stocksUseCase struct {
	stocksRepo repository.StocksRepository
}

func (s *stocksUseCase) GetStocksByName(stocksName string) (models.Stocks, error) {
	return s.stocksRepo.GetByName(stocksName)
}
func (s *stocksUseCase) UpdateByName(stocks *models.Stocks) error {
	return s.stocksRepo.Update(stocks)
}
func (s *stocksUseCase) GetAll() ([]models.Stocks, error) {
	return s.stocksRepo.GetAll()
}
func (s *stocksUseCase) GetStockQtyById(stockId string) (int, error) {
	return s.stocksRepo.GetStockQty(stockId)
}
func (s *stocksUseCase) UpdateQtyStockById(quantity int, id string) error {
	return s.stocksRepo.UpdateQtyStock(quantity, id)
}

func NewStocksUseCase(sRepo repository.StocksRepository) StocksUseCase {
	return &stocksUseCase{
		stocksRepo: sRepo,
	}
}

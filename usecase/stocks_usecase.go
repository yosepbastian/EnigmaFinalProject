package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type StocksUseCase interface {
	GetStocksByName(stocksName string) (models.Stocks, error)
	UpdateStocksById(stocks models.Stocks) error
}

type stocksUseCase struct {
	stocksRepo repository.StocksRepository
}

func (s *stocksUseCase) GetStocksByName(stocksName string) (models.Stocks, error) {
	return s.stocksRepo.GetByName(stocksName)
}
func (s *stocksUseCase) UpdateStocksById(stocks models.Stocks) error {
	return s.stocksRepo.UpdateById(stocks)
}

func NewStocksUseCase(sRepo repository.StocksRepository) StocksUseCase {
	return &stocksUseCase{
		stocksRepo: sRepo,
	}
}

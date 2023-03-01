package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type StocksUseCase interface {
	GetStocksByName(stocksName string) (models.Stocks, error)
	UpdateByName(stocks *models.Stocks) error
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

func NewStocksUseCase(sRepo repository.StocksRepository) StocksUseCase {
	return &stocksUseCase{
		stocksRepo: sRepo,
	}
}

package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)


type OrderUseCase interface{
	// CreateNewOrderBuy(newBuy *models.Stocks) error
	// CreateNewOrderSell(newSell *models.Stocks)
	GetStockByName(name string) (models.Stocks, error)

}

type orderUseCase struct{
	orderRepo repository.OrderRepository
}

func (s *orderUseCase)	GetStockByName(name string) (models.Stocks, error){
	return s.orderRepo.GetByName(name)
}

func NewOrderUseCase(orderRepoSitory repository.OrderRepository) OrderUseCase{
	return &orderUseCase{
		orderRepo: orderRepoSitory,
	}
}

package usecase

import (
	"kel1-stockbite-projects/models"
)


type OrderUseCase interface{
	CreateNewOrderBuy(newBuy *models.Stocks) error
	CreateNewOrderSell(newSell *models.Stocks)
}
package usecase

import (
	"errors"
	"fmt"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
	"strconv"
)

type OrderUseCase interface {
	// CreateNewOrderBuy(newBuy models.Stocks) error
	CreateNewOrderSell(newSell models.Transaction) error
	GetStockByName(name string) (models.Stocks, error)
}

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func (s *orderUseCase) CreateNewOrderSell(newSell models.Transaction) error {

	var tempProvit float64

	stockId := strconv.Itoa(newSell.StockID)

	stockAvailable, isStockAvailable := s.orderRepo.CheckQuantityStockUser(newSell.UserID, stockId)

	currentQuantity := stockAvailable - int (newSell.Quantity)

	if isStockAvailable != nil {

		return errors.New("you don't have this stock")

	} else if stockAvailable < int (newSell.Quantity ){

		return errors.New("your new order sell is geater than your available stocks")

	} else if newSell.Quantity < 1 {
		return errors.New("minimal quantity is 1")
	} else if int (newSell.Quantity) >  stockAvailable {
		return errors.New("maximum order sell is equal or less then your stocks")
	} else {
		tempProvit = (newSell.Price * 100) * float64(newSell.Quantity)

	}

	if currentQuantity == 0 {

		s.orderRepo.DeleteStockUser(newSell.UserID, stockId)
	} else {
		s.orderRepo.UpdateStockUser(currentQuantity, newSell.UserID, stockId)
	}


	balance, err := s.orderRepo.GetUserBalance(newSell.UserID)


	if err != nil {

		return errors.New("error has occurred when trying to get user balance")

	} else {
		newBalance := tempProvit + float64(balance)
		s.orderRepo.UpdateUserBalance(int(newBalance), newSell.UserID)

	}

	stockQuantity, Qerr := s.orderRepo.GetStockQuantityByID(stockId)
	

	if Qerr != nil {
		return errors.New("error has occurred when trying to get stock quantity")
	} else {

		newStockQuantity := stockQuantity + int ( (newSell.Quantity * 100))
		err := s.orderRepo.UpdateQuantityStock(newStockQuantity, stockId)

		if err != nil {
			return errors.New("error has occurred when trying to update stock quantity")
		}
	}

	Terr := s.orderRepo.AddNewTransaction(newSell)

	fmt.Println(Terr)


	if Terr != nil {
		return errors.New("error has occurred when trying to add new transaction")
	}

	return nil
}

func (s *orderUseCase) GetStockByName(name string) (models.Stocks, error) {
	return s.orderRepo.GetByName(name)
}

func NewOrderUseCase(orderRepoSitory repository.OrderRepository) OrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepoSitory,
	}
}

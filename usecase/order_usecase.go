package usecase

import (
	"errors"
	"fmt"
	"kel1-stockbite-projects/invoice"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
	"kel1-stockbite-projects/utils"
	"strconv"
)

type OrderUseCase interface {
	// CreateNewOrderBuy(newBuy models.Stocks) error
	OrderSell(newSell *models.Transaction) error
	OrderBuy(userId string, email string, stockName string, quantity float64, price float64) error
}

type orderUseCase struct {
	portRepo  repository.PortFoliosRepository
	stockRepo repository.StocksRepository
	txRepo    repository.TransactionRepository
	userRepo  repository.UsersRepository
}

func (s *orderUseCase) OrderSell(newSell *models.Transaction) error {

	TxnewSell := &models.Transaction{
		Id:              utils.GenerateId(),
		UserID:          newSell.UserID,
		StockID:         newSell.StockID,
		Quantity:        newSell.Quantity,
		Price:           newSell.Price * 100,
		TransactionType: "SELL",
	}

	var tempProvit float64
	stockId := strconv.Itoa(newSell.StockID)
	stockAvailable, isStockAvailable := s.portRepo.CheckQtyStock(newSell.UserID, stockId)
	currentQuantity := stockAvailable - int(newSell.Quantity)

	if isStockAvailable != nil {
		return errors.New("you don't have this stock")
	} else if stockAvailable < int(newSell.Quantity) {
		return errors.New("your new order sell is geater than your available stocks")
	} else if newSell.Quantity < 1 {
		return errors.New("minimal quantity is 1")
	} else if int(newSell.Quantity) > stockAvailable {
		return errors.New("maximum order sell is equal or less then your stocks")
	} else {
		tempProvit = (newSell.Price * 100) * float64(newSell.Quantity)

	}
	if currentQuantity == 0 {
		s.portRepo.Delete(newSell.UserID, stockId)
	} else {
		s.portRepo.UpdatePortoStok(currentQuantity, newSell.UserID, stockId)
	}
	balance, err := s.userRepo.GetUserBalance(newSell.UserID)

	if err != nil {
		return errors.New("error has occurred when trying to get user balance")
	} else {
		newBalance := tempProvit + float64(balance)
		s.userRepo.UpdateUserBalance(int(newBalance), newSell.UserID)
	}

	stockQuantity, Qerr := s.stockRepo.GetStockQty(stockId)

	if Qerr != nil {
		return errors.New("error has occurred when trying to get stock quantity")
	} else {

		newStockQuantity := stockQuantity + int((newSell.Quantity * 100))
		err := s.stockRepo.UpdateQtyStock(newStockQuantity, stockId)

		if err != nil {
			return errors.New("error has occurred when trying to update stock quantity")
		}
	}

	fmt.Println("newsell tx", TxnewSell)

	Terr := s.txRepo.Insert(TxnewSell)

	if Terr != nil {
		return errors.New("error has occurred when trying to add new transaction")
	}
	
	return nil
}

func (t *orderUseCase) OrderBuy(userId string, email string, stockName string, quantity float64, price float64) error {
	user, err := t.userRepo.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("email not registered")
	}
	//check stockname
	stock, err := t.stockRepo.GetByName(stockName)
	if err != nil {
		return fmt.Errorf("incorrect stockName")
	}
	//check price equal to stockprice
	if price != stock.Price {
		return fmt.Errorf("incorrect stock price")
	}

	totalCost := (price * 100) * quantity

	if user.Balance < totalCost {
		return fmt.Errorf("balance not enough")
	}

	floatQty := float64(quantity)

	if stock.Quantity < floatQty {
		return fmt.Errorf("stock quantity not enough")
	}
	user.Balance = user.Balance - totalCost
	stock.Quantity = stock.Quantity - floatQty

	if err := t.userRepo.Update(&user); err != nil {
		return err
	}
	if err := t.stockRepo.Update(&stock); err != nil {
		return err
	}

	transaction := models.Transaction{
		Id:              utils.GenerateId(),
		UserID:          userId,
		StockID:         stock.Id,
		Quantity:        quantity,
		Price:           price,
		TransactionType: "buy",
	}

	if err := t.txRepo.Insert(&transaction); err != nil {
		fmt.Println(err)
		return err
	}
	if err := t.portRepo.CheckAndCreate(userId, stock.Id, quantity); err != nil {
		fmt.Println(err)
		return err

	}
	err = invoice.GenerateInvoicePDF(userId, stockName, quantity, price, totalCost)
	if err != nil {
		return err
	}

	return nil
}

func NewOrderUseCase(portRepo repository.PortFoliosRepository, stockRepo repository.StocksRepository, txRepo repository.TransactionRepository, userRepo repository.UsersRepository) OrderUseCase {
	return &orderUseCase{
		portRepo:  portRepo,
		stockRepo: stockRepo,
		txRepo:    txRepo,
		userRepo:  userRepo,
	}
}

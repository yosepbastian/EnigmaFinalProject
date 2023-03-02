package usecase

import (
	"fmt"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"

	"github.com/google/uuid"
)

type TransactionUseCase interface {
	InsertNewTransaction(tx *models.Transaction) error
	BuyStocks(userId string, email string, stockName string, quantity float64, price float64) error
	SellStocks(userId string, email string, stockName string, quantity float64, price float64) error
}

type transactionUseCase struct {
	txRepo     repository.TransactionRepository
	stocksRepo repository.StocksRepository
	usersRepo  repository.UsersRepository
	portRepo   repository.PortFoliosRepository
}

func (t *transactionUseCase) InsertNewTransaction(tx *models.Transaction) error {
	return t.txRepo.Insert(tx)
}

func (t *transactionUseCase) BuyStocks(userId string, email string, stockName string, quantity float64, price float64) error {
	user, err := t.usersRepo.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("Email not registered")
	}
	//check stockname
	stock, err := t.stocksRepo.GetByName(stockName)
	if err != nil {
		return fmt.Errorf("Incorrect StockName")
	}
	//check price equal to stockprice
	if price != stock.Price {
		return fmt.Errorf("Incorrect Stock Price")
	}

	totalCost := price * quantity
	totalCostFloat := float64(totalCost)
	if user.Balance < totalCostFloat {
		return fmt.Errorf("balance not enough")
	}

	floatQty := float64(quantity)

	if stock.Quantity < floatQty {
		return fmt.Errorf("Stock Quantity Not enough")
	}
	user.Balance = user.Balance - totalCostFloat
	stock.Quantity = stock.Quantity - floatQty

	if err := t.usersRepo.Update(&user); err != nil {
		return err
	}
	if err := t.stocksRepo.Update(&stock); err != nil {
		return err
	}
	uuid := uuid.New().String()
	transaction := &models.Transaction{
		Id:              uuid,
		UserID:          userId,
		StockID:         stock.Id,
		Quantity:        quantity,
		Price:           totalCostFloat,
		TransactionType: "buy",
	}

	if err := t.InsertNewTransaction(transaction); err != nil {
		fmt.Println(err)
		return err
	}
	if err := t.portRepo.CheckAndCreate(userId, stock.Id, quantity); err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}
func (t *transactionUseCase) SellStocks(userId string, email string, stockName string, quantity float64, price float64) error {
	user, err := t.usersRepo.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("Email not registered")
	}
	//check stockname
	stock, err := t.stocksRepo.GetByName(stockName)
	if err != nil {
		return fmt.Errorf("Incorrect StockName")
	}
	//check price equal to stockprice
	if price != stock.Price {
		return fmt.Errorf("Incorrect Stock Price")
	}

	totalCost := price * quantity
	totalCostFloat := float64(totalCost)
	if user.Balance < totalCostFloat {
		return fmt.Errorf("balance not enough")
	}

	floatQty := float64(quantity)

	if stock.Quantity < floatQty {
		return fmt.Errorf("Stock Quantity Not enough")
	}
	user.Balance = user.Balance - totalCostFloat
	stock.Quantity = stock.Quantity - floatQty

	if err := t.usersRepo.Update(&user); err != nil {
		return err
	}
	if err := t.stocksRepo.Update(&stock); err != nil {
		return err
	}
	uuid := uuid.New().String()
	transaction := &models.Transaction{
		Id:              uuid,
		UserID:          userId,
		StockID:         stock.Id,
		Quantity:        quantity,
		Price:           totalCostFloat,
		TransactionType: "buy",
	}

	if err := t.InsertNewTransaction(transaction); err != nil {
		fmt.Println(err)
		return err
	}
	if err := t.portRepo.CheckAndCreate(userId, stock.Id, quantity); err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}

func NewTransactionUsecase(txRepo repository.TransactionRepository, stocksRepo repository.StocksRepository, usersRepo repository.UsersRepository, portRepo repository.PortFoliosRepository) TransactionUseCase {
	return &transactionUseCase{
		txRepo:     txRepo,
		stocksRepo: stocksRepo,
		usersRepo:  usersRepo,
		portRepo:   portRepo,
	}
}

package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type OrderRepository interface {
	GetByName(name string) (models.Stocks, error)
	CheckQuantityStockUser(userId string, stockId string) (int, error)
	DeleteStockUser(userId string, stockId string) error
	UpdateUserBalance(balance int, userId string) error
	UpdateQuantityStock(quantity int, id string) error
	UpdateStockUser(quantity int, userId string, stockId string) error
	AddNewTransaction(transaction models.Transaction) error
	GetUserBalance(userId string) (float64, error)
	GetStockPriceById(stockId string) (int, error)
	GetStockQuantityByID(stockId string) (int, error)
}

type orderRepository struct {
	db *sqlx.DB
}

func (s *orderRepository) GetByName(name string) (models.Stocks, error) {
	var stock models.Stocks

	err := s.db.Get(&stock, utils.SELECT_STOCK_NAME, name)
	if err != nil {
		return stock, err
	}

	return stock, nil
}


func (s *orderRepository) 	GetStockQuantityByID(stockId string) (int, error) {

	var quantity int

	err := s.db.Get(&quantity, utils.GET_STOCK_QUANTITY_BY_ID, stockId)

	if err != nil {
		return 0, err
	}

	return quantity, nil

}

func (s *orderRepository) 	GetStockPriceById(stockId string) (int, error) {
	var price int


	err := s.db.Get(&price, utils.GET_STOCK_PRICE_BY_ID, stockId)

	if err != nil {
		return 0, err
	}

	return price, nil
}

func (s *orderRepository) GetUserBalance(userId string) (float64, error) {
	var balance float64
	err := s.db.Get(&balance, utils.GET_USER_BALANCE, userId)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (s *orderRepository) CheckQuantityStockUser(userId string, stockId string) (int, error) {
	var quantity int

	err := s.db.Get(&quantity, utils.SELECT_QUANTITY_STOCK_USER, userId, stockId)

	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func (s *orderRepository) DeleteStockUser(userId string, stockId string) error {

	_, err := s.db.Exec(utils.DELETE_STOCK_USER, userId, stockId)

	if err != nil {
		return err
	}

	return nil
}

func (s *orderRepository) UpdateUserBalance(balance int, userId string) error {
	_, err := s.db.Exec(utils.UPDATE_USER_BALANCE, balance, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *orderRepository) UpdateQuantityStock(quantity int, id string) error {
	_, err := s.db.Exec(utils.UPDATE_QUANTITY_STOCK, quantity, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *orderRepository) UpdateStockUser(quantity int, userId string, stockId string) error {
	_, err := s.db.Exec(utils.UPDATE_QUANTITY_STOCK_USER, quantity, userId, stockId)

	if err != nil {
		return err
	}

	return nil
}

func (s *orderRepository) AddNewTransaction(transaction models.Transaction) error {

	_, err := s.db.Exec(utils.INSERT_NEW_TRANSACTION, transaction.UserID, transaction.StockID, transaction.Quantity, transaction.Price, transaction.TransactionType, transaction.Id)

	if err != nil {
		return err
	}

	return nil
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

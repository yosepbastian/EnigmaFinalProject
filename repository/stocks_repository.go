package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type StocksRepository interface {
	GetByName(name string) (models.Stocks, error)
	Update(stocks *models.Stocks) error
	GetAll() ([]models.Stocks, error)
	GetStockQty(stockId string) (int, error)
	UpdateQtyStock(quantity int, id string) error
}

type stocksRepository struct {
	db *sqlx.DB
}

func (s *stocksRepository) GetByName(name string) (models.Stocks, error) {
	var stocks models.Stocks
	err := s.db.QueryRow(utils.SELECT_STOCKS_BY_NAME, name).Scan(
		&stocks.Id,
		&stocks.Name,
		&stocks.Price,
		&stocks.Quantity,
	)
	if err != nil {
		return models.Stocks{}, err
	}
	return stocks, nil
}

func (s *stocksRepository) GetStockQty(stockId string) (int, error) {
	var quantity int
	err := s.db.Get(&quantity, utils.GET_STOCK_QUANTITY_BY_ID, stockId)
	if err != nil {
		return 0, err
	}
	return quantity, nil

}

func (s *stocksRepository) Update(stocks *models.Stocks) error {
	_, err := s.db.NamedExec(utils.UPDATE_STOCKS, stocks)
	if err != nil {
		return err
	}
	return nil
}
func (s *stocksRepository) GetAll() ([]models.Stocks, error) {
	var stocks []models.Stocks
	err := s.db.Select(&stocks, utils.GetAll)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

func (s *stocksRepository) UpdateQtyStock(quantity int, id string) error {
	_, err := s.db.Exec(utils.UPDATE_QUANTITY_STOCK, quantity, id)

	if err != nil {
		return err
	}

	return nil
}

func NewStocksRepository(db *sqlx.DB) StocksRepository {
	return &stocksRepository{
		db: db,
	}
}

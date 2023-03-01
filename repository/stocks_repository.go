package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type StocksRepository interface {
	GetByName(name string) (models.Stocks, error)
	Update(stocks *models.Stocks) error
	
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

func (s *stocksRepository) Update(stocks *models.Stocks) error {
	_, err := s.db.NamedExec(utils.UPDATE_STOCKS, stocks)
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

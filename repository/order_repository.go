package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type OrderRepository interface{
	GetByName(name string) (models.Stocks, error)
}

type orderRepository struct{
	db *sqlx.DB
}


func (s *orderRepository) GetByName(name string) ( models.Stocks, error){
	var stock models.Stocks

	err := s.db.QueryRow(utils.SELECT_STOCK_NAME, name).Scan(
		&stock.Id,
		&stock.Name,
		&stock.Price,
		&stock.Quantity,
	)

	if err != nil {
		return models.Stocks{}, err
	}

	return models.Stocks{},nil
}

func NewOrderRepository(db *sqlx.DB) OrderRepository{
	return &orderRepository{
		db: db,
	}
}
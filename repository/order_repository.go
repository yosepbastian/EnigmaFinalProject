package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type OrderRepository interface {
	GetByName(name string) (models.Stocks, error)
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

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

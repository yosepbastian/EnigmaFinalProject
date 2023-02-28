package repository

import (
	"kel1-stockbite-projects/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(transaksi *models.Transaction)
}

type TransactionRepository struct {
	db *sqlx.DB
}

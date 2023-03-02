package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(tx models.Transaction) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func (t *transactionRepository) Insert(tx models.Transaction) error {
	_, err := t.db.NamedExec(utils.INSERT_TRANSACTION, tx)
	if err != nil {
		return err
	}
	return nil
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

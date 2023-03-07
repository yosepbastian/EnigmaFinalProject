package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(tx *models.Transaction) error
	GetAll() ([]models.TransactionAdm, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func (t *transactionRepository) Insert(tx *models.Transaction) error {
	_, err := t.db.NamedExec(utils.INSERT_TRANSACTION, tx)
	if err != nil {
		return err
	}
	return nil
}
func (t *transactionRepository) GetAll() ([]models.TransactionAdm, error) {
	var tx []models.TransactionAdm
	err := t.db.Select(&tx, utils.GetAll_TRANSACTION)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

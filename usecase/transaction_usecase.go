package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type TransactionUseCase interface {
	InsertNewTransaction(tx models.Transaction) error
	GetAllTrancaction() ([]models.TransactionAdm, error)
}

type transactionUseCase struct {
	txRepo     repository.TransactionRepository
	stocksRepo repository.StocksRepository
	usersRepo  repository.UsersRepository
	portRepo   repository.PortFoliosRepository
}

func (t *transactionUseCase) InsertNewTransaction(tx models.Transaction) error {
	return t.txRepo.Insert(&tx)
}
func (t *transactionUseCase) GetAllTrancaction() ([]models.TransactionAdm, error) {
	return t.txRepo.GetAll()
}

func NewTransactionUsecase(txRepo repository.TransactionRepository, stocksRepo repository.StocksRepository, usersRepo repository.UsersRepository, portRepo repository.PortFoliosRepository) TransactionUseCase {
	return &transactionUseCase{
		txRepo:     txRepo,
		stocksRepo: stocksRepo,
		usersRepo:  usersRepo,
		portRepo:   portRepo,
	}
}

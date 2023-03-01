package manager

import "kel1-stockbite-projects/usecase"

type UseCaseManager interface {
	StocksUseCase() usecase.StocksUseCase
	BuyStocks() usecase.TransactionUseCase
	OrderUseCase() usecase.OrderUseCase
}

type usecaseManager struct {
	repomanager RepositoryManager
}

func (u *usecaseManager) StocksUseCase() usecase.StocksUseCase {
	return usecase.NewStocksUseCase(u.repomanager.StocksRepository())
}

func (u *usecaseManager) OrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderUseCase(u.repomanager.OrderRepository())
}


func (u *usecaseManager) BuyStocks() usecase.TransactionUseCase {
	return usecase.NewTransactionUsecase(
		u.repomanager.TransactionRepository(),
		u.repomanager.StocksRepository(),
		u.repomanager.UsersRepository(),
		u.repomanager.PortfoliosRepository(),
	)
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &usecaseManager{
		repomanager: repoManager,
	}
}

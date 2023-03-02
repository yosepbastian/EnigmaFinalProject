package manager

import "kel1-stockbite-projects/usecase"

type UseCaseManager interface {
	StocksUseCase() usecase.StocksUseCase
	OrderUseCase() usecase.OrderUseCase
	
}

type usecaseManager struct {
	repomanager RepositoryManager
}

func (u *usecaseManager) StocksUseCase() usecase.StocksUseCase {
	return usecase.NewStocksUseCase(u.repomanager.StocksRepository())
}
func (u *usecaseManager) UserUseCase() usecase.UsersUseCase {
	return usecase.NewUsersUseCase(u.repomanager.UsersRepository())
}

func (u *usecaseManager) OrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderUseCase(u.repomanager.PortfoliosRepository(), u.repomanager.StocksRepository(), u.repomanager.TransactionRepository(), u.repomanager.UsersRepository())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &usecaseManager{
		repomanager: repoManager,
	}
}

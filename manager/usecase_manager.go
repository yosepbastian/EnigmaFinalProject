package manager

import "kel1-stockbite-projects/usecase"

type UseCaseManager interface {
	StocksUseCase() usecase.StocksUseCase
}

type usecaseManager struct {
	repomanager RepositoryManager
}

func (u *usecaseManager) StocksUseCase() usecase.StocksUseCase {
	return usecase.NewStocksUseCase(u.repomanager.StocksUseCase())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &usecaseManager{
		repomanager: repoManager,
	}
}

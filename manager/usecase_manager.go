package manager

import "kel1-stockbite-projects/usecase"

type UseCaseManager interface{
	OrderUseCase() usecase.OrderUseCase
}

	type usecaseManager struct {
		repomanager RepositoryManager
	}

	func (u *usecaseManager) OrderUseCase() usecase.OrderUseCase {
		return usecase.NewOrderUseCase(u.repomanager.OrderRepository())
	}


	func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager{
		return &usecaseManager{
			repomanager: repoManager,
		}
	}


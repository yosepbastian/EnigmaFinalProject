package manager

import "kel1-stockbite-projects/repository"


type RepositoryManager interface{
	OrderRepository() repository.OrderRepository
}

type repositoryManager struct{

	infra InfraManager
}

func(r *repositoryManager) OrderRepository() repository.OrderRepository{

	return repository.NewOrderRepository(r.infra.SqlDb())
}


func NewRepositoryManager(infraManager InfraManager) RepositoryManager{
	return &repositoryManager{
		infra: infraManager,
	}
}

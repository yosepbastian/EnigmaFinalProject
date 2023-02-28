package manager

import "kel1-stockbite-projects/repository"


type RepositoryManager interface{
	StocksUseCase() repository.StocksRepository
}

type repositoryManager struct{

	infra InfraManager
}

func(r *repositoryManager) StocksUseCase() repository.StocksRepository{

	return repository.NewStocksRepository(r.infra.SqlDb())
}


func NewRepositoryManager(infraManager InfraManager) RepositoryManager{
	return &repositoryManager{
		infra: infraManager,
	}
}

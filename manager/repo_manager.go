package manager

import "kel1-stockbite-projects/repository"

type RepositoryManager interface {
	UsersRepository() repository.UsersRepository
	PortfoliosRepository() repository.PortFoliosRepository
	TransactionRepository() repository.TransactionRepository
	StocksRepository() repository.StocksRepository

}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) UsersRepository() repository.UsersRepository {
	return repository.NewUsersRepository(r.infra.SqlDb())
}
func (r *repositoryManager) PortfoliosRepository() repository.PortFoliosRepository {
	return repository.NewPortFoliosRepository(r.infra.SqlDb())
}
func (r *repositoryManager) TransactionRepository() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.SqlDb())
}
func (r *repositoryManager) StocksRepository() repository.StocksRepository {
	return repository.NewStocksRepository(r.infra.SqlDb())
}




func NewRepositoryManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infraManager,
	}
}

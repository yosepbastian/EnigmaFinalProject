package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type PortfoliosUseCase interface {
	InsertPortfolios(porto *models.PortFolios) error
	UpdatePorto(porto *models.PortFolios) error
	GetByUserID(userID string, stockId int) (models.PortFolios, error)
	CheckQtyStockUser(userId string, stockId string) (int, error)
	CheckAndCreatePortUser(userId string, stockId int, quantity float64) error
}

type portfoliosUseCase struct {
	portfoliosRepo repository.PortFoliosRepository
}

func (p *portfoliosUseCase) InsertPortfolios(porto *models.PortFolios) error {
	return p.portfoliosRepo.Insert(porto)
}
func (p *portfoliosUseCase) GetByUserID(userID string, stockId int) (models.PortFolios, error) {
	return p.portfoliosRepo.GetByIdandStockId(userID, stockId)
}
func (p *portfoliosUseCase) UpdatePorto(porto *models.PortFolios) error {
	return p.portfoliosRepo.Update(porto)
}
func (p *portfoliosUseCase) CheckQtyStockUser(userId string, stockId string) (int, error) {
	return p.portfoliosRepo.CheckQtyStock(userId, stockId)
}
func (p *portfoliosUseCase) CheckAndCreatePortUser(userId string, stockId int, quantity float64) error {
	return p.portfoliosRepo.CheckAndCreate(userId, stockId, quantity)
}

func NewPortfoliosUseCase(pRepo repository.PortFoliosRepository) PortfoliosUseCase {
	return &portfoliosUseCase{
		portfoliosRepo: pRepo,
	}
}

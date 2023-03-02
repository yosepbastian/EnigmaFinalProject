package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type PortfoliosUseCase interface {
	InsertPortfolios(porto *models.PortFolios) error
	GetPortoByIdandStockId(userID string, stockID int) (models.PortFolios, error)
	Update(porto *models.PortFolios) error
	Delete(userId string, stockId string) error
	CheckAndCreatePortUser(userId string, stockId int, quantity float64) error
	CheckQtyStockUser(userId string, stockId string) (int, error)
	UpdatePortoStok(quantity int, userId string, stockId string) error
}

type portfoliosUseCase struct {
	portfoliosRepo repository.PortFoliosRepository
}

func (p *portfoliosUseCase) InsertPortfolios(porto *models.PortFolios) error {
	return p.portfoliosRepo.Insert(porto)
}

func (p *portfoliosUseCase) GetPortoByIdandStockId(userID string, stockID int) (models.PortFolios, error) {
	return p.portfoliosRepo.GetByIdandStockId(userID, stockID)
}

func (p *portfoliosUseCase) Delete(userId string, stockId string) error {
	return p.portfoliosRepo.Delete(userId, stockId)
}
func (p *portfoliosUseCase) Update(porto *models.PortFolios) error {
	return p.portfoliosRepo.Update(porto)
}

func (p *portfoliosUseCase) CheckAndCreatePortUser(userId string, stockId int, quantity float64) error {
	return p.portfoliosRepo.CheckAndCreate(userId, stockId, quantity)
}

func (p *portfoliosUseCase) CheckQtyStockUser(userId string, stockId string) (int, error) {
	return p.portfoliosRepo.CheckQtyStock(userId, stockId)
}
func (p *portfoliosUseCase) UpdatePortoStok(quantity int, userId string, stockId string) error {
	return p.portfoliosRepo.UpdatePortoStok(quantity, userId, stockId)
}

func NewPortfoliosUseCase(pRepo repository.PortFoliosRepository) PortfoliosUseCase {
	return &portfoliosUseCase{
		portfoliosRepo: pRepo,
	}
}

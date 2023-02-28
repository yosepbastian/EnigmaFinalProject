package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type PortfoliosUseCase interface {
	InsertPortfolios(tx *models.PortFolios) error
	Update(p *models.PortFolios) error
	GetByUserIDAndStockID(userID string, stockID int) (*models.PortFolios, error)
}

type portfoliosUseCase struct {
	portfoliosRepo repository.PortFoliosRepository
}

func (p *portfoliosUseCase) InsertPortfolios(tx *models.PortFolios) error {
	return p.portfoliosRepo.Insert(tx)
}
func (p *portfoliosUseCase) GetByUserIDAndStockID(userID string, stockID int) (*models.PortFolios, error) {
	return p.portfoliosRepo.GetByUserIDAndStockID(userID, stockID)
}
func (p *portfoliosUseCase) Update(pr *models.PortFolios) error {
	return p.portfoliosRepo.Update(pr)
}

func NewPortfoliosUseCase(pRepo repository.PortFoliosRepository) PortfoliosUseCase {
	return &portfoliosUseCase{
		portfoliosRepo: pRepo,
	}
}

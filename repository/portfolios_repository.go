package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type PortFoliosRepository interface {
	Insert(tx *models.PortFolios) error
	GetByIdandStockId(userID string, stockID int) (models.PortFolios, error)
	Update(p *models.PortFolios) error
}

type portFoliosRepository struct {
	db *sqlx.DB
}

func (p *portFoliosRepository) Insert(tx *models.PortFolios) error {
	_, err := p.db.NamedExec(utils.INSERT_PORTFOLIOS, tx)

	if err != nil {
		panic(err.Error())
	}
	return nil
}

func (r *portFoliosRepository) Update(p *models.PortFolios) error {
	_, err := r.db.NamedExec(utils.UPDATE_PORTFOLIOS, p)
	if err != nil {
		return err
	}
	return nil
}

func (p *portFoliosRepository) GetByIdandStockId(userID string, stockID int) (models.PortFolios, error) {
	var portfolio models.PortFolios
	err := p.db.QueryRow(utils.GET_BY_USER_ID_AND_STOCK_ID, userID, stockID).Scan(
		&portfolio.Id,
		&portfolio.UserID,
		&portfolio.StockID,
		&portfolio.Quantity,
	)
	if err != nil {
		return models.PortFolios{}, err
	}
	return portfolio, nil
}

func NewPortFoliosRepository(db *sqlx.DB) PortFoliosRepository {
	return &portFoliosRepository{
		db: db,
	}
}

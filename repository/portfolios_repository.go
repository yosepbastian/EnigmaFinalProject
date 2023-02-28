package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type PortFoliosRepository interface {
	Insert(tx *models.PortFolios) error
	GetByUserIDAndStockID(userID string, stockID int) (*models.PortFolios, error)
	Update(p *models.PortFolios) error
}

func (r *portFoliosRepository) Update(p *models.PortFolios) error {

	_, err := r.db.Exec(utils.UPDATE_PORTFOLIOS, p.Quantity, p.UserID, p.StockID)
	if err != nil {
		return err
	}
	return nil
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

func (p *portFoliosRepository) GetByUserIDAndStockID(userID string, stockID int) (*models.PortFolios, error) {
	var portfolio models.PortFolios

	err := p.db.Get(&portfolio, utils.GET_BY_USERID_AND_STOCKID, userID, stockID)
	if err != nil {
		panic(err.Error())
	}

	return &portfolio, nil
}

func NewPortFoliosRepository(db *sqlx.DB) PortFoliosRepository {
	return &portFoliosRepository{
		db: db,
	}
}

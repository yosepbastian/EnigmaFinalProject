package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PortFoliosRepository interface {
	Insert(porto *models.PortFolios) error
	GetByIdandStockId(userID string, stockID int) (models.PortFolios, error)
	Update(porto *models.PortFolios) error
	Delete(userId string, stockId string) error
	CheckAndCreate(userId string, stockId int, quantity float64) error
	CheckQtyStock(userId string, stockId string) (int, error)
	UpdatePortoStok(quantity int, userId string, stockId string) error
}

type portFoliosRepository struct {
	db *sqlx.DB
}

func (p *portFoliosRepository) Insert(porto *models.PortFolios) error {
	_, err := p.db.NamedExec(utils.INSERT_PORTFOLIOS, porto)

	if err != nil {
		panic(err.Error())
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

func (r *portFoliosRepository) Update(porto *models.PortFolios) error {
	_, err := r.db.NamedExec(utils.UPDATE_PORTFOLIOS, porto)
	if err != nil {
		return err
	}
	return nil
}

func (s *portFoliosRepository) Delete(userId string, stockId string) error {

	_, err := s.db.Exec(utils.DELETE_STOCK_USER, userId, stockId)

	if err != nil {
		return err
	}

	return nil
}

func (t *portFoliosRepository) CheckAndCreate(userId string, stockId int, quantity float64) error {
	portfolio, err := t.GetByIdandStockId(userId, stockId)
	if err != nil {
		uuid := uuid.New().String()
		// If portfolio entry does not exist, create a new one
		portfolio = models.PortFolios{
			Id:       uuid,
			UserID:   userId,
			StockID:  stockId,
			Quantity: quantity,
		}
		if err := t.Insert(&portfolio); err != nil {
			return err
		}
	} else {
		// If portfolio entry already exists, update the quantity
		portfolio.Quantity = portfolio.Quantity + quantity
		if err := t.Update(&portfolio); err != nil {
			return err
		}
	}

	return nil
}

func (s *portFoliosRepository) CheckQtyStock(userId string, stockId string) (int, error) {
	var quantity int
	err := s.db.Get(&quantity, utils.SELECT_QUANTITY_STOCK_USER, userId, stockId)
	if err != nil {
		return 0, err
	}
	return quantity, nil
}

func (s *portFoliosRepository) UpdatePortoStok(quantity int, userId string, stockId string) error {
	_, err := s.db.Exec(utils.UPDATE_QUANTITY_STOCK_USER, quantity, userId, stockId)

	if err != nil {
		return err
	}

	return nil
}

func NewPortFoliosRepository(db *sqlx.DB) PortFoliosRepository {
	return &portFoliosRepository{
		db: db,
	}
}

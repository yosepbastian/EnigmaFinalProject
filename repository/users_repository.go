package repository

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type UsersRepository interface {
	Insert(newUser *models.Users) error
}

type usersRepository struct {
	db *sqlx.DB
}

func (u *usersRepository) Insert(newUser *models.Users) error {
	_, err := u.db.NamedExec(utils.INSERT_USER, newUser)
	if err != nil {
		return err
	}
	return nil
}

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

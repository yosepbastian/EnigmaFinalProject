package repository

import (
	"database/sql"
	"fmt"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils"

	"github.com/jmoiron/sqlx"
)

type UsersRepository interface {
	Insert(newUser *models.Users) error
	GetById(id string) (models.Users, error)
	Update(id *models.Users) error
	GetByEmail(email string) (models.Users, error)
	Begin() (tx *sql.Tx, err error)
	GetByEmailForUpdate(email string, tx *sql.Tx) (*models.Users, error)
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
func (t *usersRepository) Begin() (tx *sql.Tx, err error) {
	tx, err = t.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	return tx, nil
}

func (u *usersRepository) GetByEmail(email string) (models.Users, error) {
	var users models.Users
	err := u.db.QueryRow(utils.SELECT_USER_BY_EMAIL, email).Scan(
		&users.Id,
		&users.Name,
		&users.Email,
		&users.Password,
		&users.Balance,
	)
	if err != nil {
		return models.Users{}, err
	}
	return users, nil
}

func (u *usersRepository) Update(id *models.Users) error {

	_, err := u.db.NamedExec(utils.UPDATE_USER, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *usersRepository) GetById(id string) (models.Users, error) {
	var users models.Users
	err := u.db.QueryRow(utils.SELECT_USER_ID, id).Scan(
		&users.Id,
		&users.Name,
		&users.Email,
		&users.Password,
		&users.Password,
	)
	if err != nil {
		return models.Users{}, err
	}
	return users, nil
}

func (u *usersRepository) GetByEmailForUpdate(email string, tx *sql.Tx) (*models.Users, error) {
	row := tx.QueryRow(utils.SELECT_EMAIL_FOR_UPDATE, email)

	user := &models.Users{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Balance)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

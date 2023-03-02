package usecase

import (
	"database/sql"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type UsersUseCase interface {
	RegisterUser(newUser *models.Users) error
	UpdateUser(user *models.Users) error
	GetByIdUser(id string) (models.Users, error)
	GetByEmailUser(email string) (models.Users, error)
	GetByEmailForUpdate(email string, tx *sql.Tx) (*models.Users, error)
	GetUserBalanceById(userId string) (float64, error)
	UpdateUserBalanceByUserId(balance int, userId string) error
}

type usersUseCase struct {
	usersRepo repository.UsersRepository
}

func (u *usersUseCase) RegisterUser(newUser *models.Users) error {
	return u.usersRepo.Insert(newUser)
}
func (u *usersUseCase) UpdateUser(user *models.Users) error {
	return u.usersRepo.Update(user)
}
func (u *usersUseCase) GetByEmailUser(email string) (models.Users, error) {
	return u.usersRepo.GetByEmail(email)
}
func (u *usersUseCase) GetByIdUser(id string) (models.Users, error) {
	return u.usersRepo.GetById(id)
}

func (u *usersUseCase) GetByEmailForUpdate(email string, tx *sql.Tx) (*models.Users, error) {
	return u.usersRepo.GetByEmailForUpdate(email, tx)
}
func (u *usersUseCase) GetUserBalanceById(userId string) (float64, error) {
	return u.usersRepo.GetUserBalance(userId)
}

func (u *usersUseCase) UpdateUserBalanceByUserId(balance int, userId string) error {
	return u.usersRepo.UpdateUserBalance(balance, userId)
}

func NewUsersUseCase(uRepo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{
		usersRepo: uRepo,
	}
}

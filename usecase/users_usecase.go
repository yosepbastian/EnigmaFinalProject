package usecase

import (
	"database/sql"
	"errors"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
	"strings"

	"github.com/google/uuid"
)

type UsersUseCase interface {
	RegisterUser(newUser *models.Users) error
	UpdateUser(user *models.Users) error
	GetByIdUser(id string) (models.Users, error)
	GetByEmailUser(email string) (models.Users, error)
	GetByEmailForUpdate(email string, tx *sql.Tx) (*models.Users, error)
	GetUserBalanceById(userId string) (float64, error)
	UpdateUserBalanceByUserId(balance int, userId string) error
	LoginUser(email string, password string) (bool, error)
	GetUserByPassword(password string) (models.Users, error)
}
type usersUseCase struct {
	usersRepo repository.UsersRepository
}

func NewUsersUseCase(uRepo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{
		usersRepo: uRepo,
	}
}

func (u *usersUseCase) RegisterUser(newUser *models.Users) error {
	uuid := uuid.New().String()
	newUser.Id = uuid

	if newUser.Email == "" || !strings.Contains(newUser.Email, "@") {
		return errors.New("invalid email")
	}

	if len(newUser.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasUpperCase := false
	hasLowerCase := false
	hasNumber := false

	for _, ch := range newUser.Password {
		if 'A' <= ch && ch <= 'Z' {
			hasUpperCase = true
		} else if 'a' <= ch && ch <= 'z' {
			hasLowerCase = true
		} else if '0' <= ch && ch <= '9' {
			hasNumber = true
		}
	}

	if !hasUpperCase || !hasLowerCase || !hasNumber {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, and one number")
	}

	return u.usersRepo.Insert(newUser)
}

func (uc *usersUseCase) LoginUser(email string, password string) (bool, error) {

	user, err := uc.GetByEmailUser(email)
	if err != nil {
		return false, errors.New("incorrect email")
	}

	// Check if password is correct
	if user.Password != password {
		return false, errors.New("incorrect password")
	}

	return true, nil
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
func (u *usersUseCase) GetUserByPassword(password string) (models.Users, error) {
	return u.usersRepo.GetByPassword(password)
}

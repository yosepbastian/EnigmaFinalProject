package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
)

type UsersUseCase interface {
	RegisterUser(newUser *models.Users) error
}

type usersUseCase struct {
	usersRepo repository.UsersRepository
}

func (u *usersUseCase) RegisterUser(newUser *models.Users) error {
	return u.usersRepo.Insert(newUser)
}

func NewUsersUseCase(uRepo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{
		usersRepo: uRepo,
	}
}

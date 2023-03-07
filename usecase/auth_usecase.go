package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/repository"
	"kel1-stockbite-projects/utils/authenticator"
)

//dummy credential just for  testing

type AuthUseCase interface {
	UserAuth(user models.UserLogin) (token string, err error)
	AdminAuth(admin models.AdminLogin) (token string, err error)
}
type authUseCase struct {
	tokenService authenticator.AccessToken

	userValidate repository.UsersRepository
}

func (a *authUseCase) UserAuth(user models.UserLogin) (token string, err error) {

	newName, _ := a.userValidate.GetUserName(user.Email)
	err, valid := a.userValidate.ValidateUserLogin(user.Email, user.Password)

	if err != nil {
		return "", err
	}

	if valid {
		token, err := a.tokenService.CreateAccessToken(user.Email, newName)
		if err != nil {
			return "nil", err
		}
		return token, nil
	} else {
		return "wrong email or password", err
	}

}
func (a *authUseCase) AdminAuth(admin models.AdminLogin) (token string, err error) {

	Name := "Admin"
	email := "stockbite@gmail.com"
	password := "kelompok1"

	if email == admin.Email && password == admin.Password {
		token, err := a.tokenService.CreateAccessToken(admin.Email, Name)
		if err != nil {
			return "nil", err
		}
		return token, nil
	} else {
		return "wrong email or password", err
	}

}

func NewAuthUseCase(service authenticator.AccessToken, userValidate repository.UsersRepository) AuthUseCase {
	authUseCase := new(authUseCase)
	authUseCase.tokenService = service
	authUseCase.userValidate = userValidate

	return authUseCase
}

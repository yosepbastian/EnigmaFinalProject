package usecase

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/utils/authenticator"
)

//dummy credential just for  testing



type AuthUseCase interface{
	UserAuth(user models.UserLogin) (token string, err error)
}
type authUseCase struct {
	tokenService authenticator.AccessToken
}

func (a *authUseCase) UserAuth(user models.UserLogin) (token string, err error) {

	// dmmyName := "Carly Cruikshank"
dmmyEmail := "ccruikshank0@slideshare.net"
dmmyPassword := "asfweg3t4vSFGQWE"


	if user.Email == dmmyEmail && user.Password == dmmyPassword {	
		token, err := a.tokenService.CreateAccessToken(&user)
		if err != nil {
			return "nil", err
		}
		return token, nil
	} else {
		return "wrong email or password", err
	}

}

func NewAuthUseCase(service authenticator.AccessToken) AuthUseCase{
	authUseCase := new(authUseCase)
	authUseCase.tokenService = service

	return authUseCase
}

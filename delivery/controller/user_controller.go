package controller

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	user usecase.UsersUseCase
}

func (oc *UserController) SignUp(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := oc.user.RegisterUser(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Register Succes",
	})
}

func (oc *UserController) Login(ctx *gin.Context) {
	var user models.UsersLogin
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Succes Login",
	})

}

func NewUserController(router *gin.Engine, userUc usecase.UsersUseCase) *UserController {
	newUserController := UserController{
		userUc,
	}
	router.POST("users/signup", newUserController.SignUp)
	router.POST("users/login", newUserController.Login)
	return &newUserController
}

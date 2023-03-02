package controller

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	user usecase.UsersUseCase
}

func (oc *UserController) RegisterUser(ctx *gin.Context) {
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

	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func NewUserController(router *gin.Engine, userUc usecase.UsersUseCase) *UserController {
	newUserController := UserController{
		userUc,
	}
	router.POST("/registerUser", newUserController.RegisterUser)
	return &newUserController
}

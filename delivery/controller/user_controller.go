package controller

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/token"
	"kel1-stockbite-projects/usecase"
	"net/http"

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

	token, err := token.GenerateToken(user.Email, user.Password)
	if err != nil {
		ctx.AbortWithStatus(401)
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})

}

func Validate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Logged in",
	})
}

func NewUserController(router *gin.Engine, userUc usecase.UsersUseCase) *UserController {
	newUserController := UserController{
		userUc,
	}
	router.GET("users/validate", Validate)
	router.POST("users/signup", newUserController.SignUp)
	router.POST("users/login", newUserController.Login)
	return &newUserController
}

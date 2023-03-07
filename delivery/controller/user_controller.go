package controller

import (
	"fmt"
	"kel1-stockbite-projects/delivery/middleware"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	rg           *gin.RouterGroup
	authUseCase  usecase.AuthUseCase
	portoUsecase usecase.PortfoliosUseCase
	userUseCase  usecase.UsersUseCase
}

func (oc *UserController) SignUpUser(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := oc.userUseCase.RegisterUser(&user)
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

func (oc *UserController) UserAuth(ctx *gin.Context) {
	var userLogin models.UserLogin

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := oc.authUseCase.UserAuth(userLogin)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": token,
	})
}

func (oc *UserController) GetPortfolio(ctx *gin.Context) {

	var userId models.PortoUserID

	err := ctx.ShouldBindJSON(&userId)

	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
	} else {
		newAsset, err := oc.portoUsecase.GetPortoByUserId(userId.UserID)
		fmt.Println(userId.UserID)
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": err.Error(),
			})
		} else {

			ctx.JSON(200, gin.H{
				"message": "OK",
				"data":    newAsset,
			})
		}
	}
}

func NewUserController(routerGroup *gin.RouterGroup, userUc usecase.UsersUseCase, authUseCase usecase.AuthUseCase, tokenMdw middleware.AuthTokenMiddleWare, portoUsecase usecase.PortfoliosUseCase) *UserController {
	newUserController := UserController{
		rg:           routerGroup,
		authUseCase:  authUseCase,
		portoUsecase: portoUsecase,
		userUseCase:  userUc,
	}
	newUserController.rg.POST("/auth", newUserController.UserAuth)
	newUserController.rg.POST("/signup", newUserController.SignUpUser)

	protectedGroup := newUserController.rg.Group("/user", tokenMdw.RequireToken())
	protectedGroup.GET("/portfolio", newUserController.GetPortfolio)

	return &newUserController
}

package controller

import (
	"kel1-stockbite-projects/delivery/middleware"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	rg          *gin.RouterGroup
	userUseCase usecase.UsersUseCase
	authUseCase usecase.AuthUseCase
	txUseCase   usecase.TransactionUseCase
}

func (oc *AdminController) AdminAuth(ctx *gin.Context) {
	var admin models.AdminLogin

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := oc.authUseCase.AdminAuth(admin)
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

func (oc *AdminController) GetAllTransaction(ctx *gin.Context) {
	tx, err := oc.txUseCase.GetAllTrancaction()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "OK",
			"data":    tx,
		})
	}
}

func NewAdminController(routerGroup *gin.RouterGroup, usersUc usecase.UsersUseCase, authUc usecase.AuthUseCase, txUc usecase.TransactionUseCase, tokenMdw middleware.AuthTokenMiddleWare) *AdminController {
	newAdminController := AdminController{
		rg:          routerGroup,
		userUseCase: usersUc,
		authUseCase: authUc,
		txUseCase:   txUc,
	}
	newAdminController.rg.POST("/admin/auth", newAdminController.AdminAuth)

	protectedGroup := newAdminController.rg.Group("/admin", tokenMdw.RequireToken())
	protectedGroup.GET("/transaction", newAdminController.GetAllTransaction)

	return &newAdminController
}

package controller

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderUsecase usecase.OrderUseCase
}

func (oc *OrderController) GetStockByName(ctx *gin.Context) {
	stock, err := oc.orderUsecase.GetStockByName("BBCA")

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "OK",
			"data":    stock,
		})
	}
}

func (oc *OrderController) CreateNewOrderSell(ctx *gin.Context) {
	var newSell models.Transaction

	err := oc.orderUsecase.CreateNewOrderSell(newSell)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "OK",
			"data":    newSell,
		})
	}
}

func NewOrderController(router *gin.Engine, orderUc usecase.OrderUseCase) *OrderController {
	newOrderController := OrderController{
		orderUsecase: orderUc,
	}

	router.GET("/stockname", newOrderController.GetStockByName)
	router.POST("/sell", newOrderController.CreateNewOrderSell)
	return &newOrderController
}

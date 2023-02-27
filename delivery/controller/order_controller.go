package controller

import (
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderUsecase usecase.OrderUseCase
}

func (oc *OrderController) GetStockByName (ctx *gin.Context) {
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

func NewOrderController(router *gin.Engine, orderUc usecase.OrderUseCase) *OrderController {
	newOrderController := OrderController{
		orderUsecase: orderUc,
	}

	router.GET("/stockname",newOrderController.GetStockByName)
	return &newOrderController
}

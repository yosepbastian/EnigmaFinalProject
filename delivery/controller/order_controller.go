package controller

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	BuysAndSell usecase.OrderUseCase
}

func (pc *OrderController) OrderBuy(ctx *gin.Context) {
	var request models.OrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := pc.BuysAndSell.OrderBuy(request.UserID, request.Email, request.StockName, request.Quantity, request.Price)
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

func (oc *OrderController) OrderSell(ctx *gin.Context) {
	var newSell models.Transaction

	if err := ctx.ShouldBindJSON(&newSell); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := oc.BuysAndSell.OrderSell(newSell)
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

func NewOrderController(router *gin.Engine, orderUc usecase.OrderUseCase) *OrderController {
	newOrderController := OrderController{
		orderUc,
	}
	router.POST("/stocks/buy", newOrderController.OrderBuy)
	router.POST("stocks/sell", newOrderController.OrderSell)
	return &newOrderController
}

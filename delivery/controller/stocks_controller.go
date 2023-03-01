package controller

import (
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type StocksController struct {
	stocks       usecase.StocksUseCase
	transactions usecase.TransactionUseCase
	sell usecase.OrderUseCase
}

func (pc *StocksController) GetById(ctx *gin.Context) {
	customer, err := pc.stocks.GetStocksByName("BBCA")
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "OK",
			"data":    customer,
		})
	}
}

func (pc *StocksController) BuyStocks(ctx *gin.Context) {
	var request models.OrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := pc.transactions.BuyStocks(request.UserID, request.Email, request.StockName, request.Quantity, request.Price)
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


func (oc *StocksController) CreateNewOrderSell(ctx *gin.Context) {
	var newSell models.Transaction

	if err := ctx.ShouldBindJSON(&newSell); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}


	err := oc.sell.CreateNewOrderSell(newSell)
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

func NewStocksController(router *gin.Engine, stocksUc usecase.StocksUseCase, transactionsUc usecase.TransactionUseCase, orderUc usecase.OrderUseCase) *StocksController {
	newStocksController := StocksController{
		stocksUc,
		transactionsUc,
		orderUc,
	}
	router.GET("/stocks/name", newStocksController.GetById)
	router.POST("/stocks/buy", newStocksController.BuyStocks)
	router.POST("stocks/sell", newStocksController.CreateNewOrderSell)
	return &newStocksController
}

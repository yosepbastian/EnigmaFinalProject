package controller

import (
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type StocksController struct {
	stocks usecase.StocksUseCase
}

func (pc *StocksController) GetAll(ctx *gin.Context) {
	stocks, err := pc.stocks.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "OK",
			"data":    stocks,
		})
	}
}

func NewStocksController(router *gin.Engine, stocksUc usecase.StocksUseCase) *StocksController {
	newStocksController := StocksController{
		stocksUc,
	}
	router.GET("/stocks", newStocksController.GetAll)
	return &newStocksController
}

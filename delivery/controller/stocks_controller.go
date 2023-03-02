package controller

import (
	"kel1-stockbite-projects/delivery/middleware"
	"kel1-stockbite-projects/models"
	"kel1-stockbite-projects/usecase"

	"github.com/gin-gonic/gin"
)

type StocksController struct {
	rg *gin.RouterGroup

	stocks       usecase.StocksUseCase
	transactions usecase.TransactionUseCase
	sell         usecase.OrderUseCase
	authUseCase  usecase.AuthUseCase
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

func (oc *StocksController) UserAuth(ctx *gin.Context) {
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

// func (pc *StocksController) UserAuth(ctx *gin.Context) {
// 	var user models.UserLogin

// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "can't bind struct",
// 		})
// 		return
// 	}

// 	token, err := pc.authUseCase.UserAuth(user)

// 	if err != nil {
// 		ctx.AbortWithStatus(401)
// 		return
// 	}

// 	ctx.JSON(200, gin.H{
// 		"token": token,
// 	})

// }

func (oc *StocksController) CreateNewOrderSell(ctx *gin.Context) {
	var newSell models.Transaction

	if err := ctx.ShouldBindJSON(newSell); err != nil {
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

func NewStocksController(routerGroup *gin.RouterGroup, stocksUc usecase.StocksUseCase, transactionsUc usecase.TransactionUseCase, orderUc usecase.OrderUseCase, authUseCase usecase.AuthUseCase, tokenMdw middleware.AuthTokenMiddleWare) *StocksController {
	newStocksController := StocksController{
		rg:           routerGroup,
		stocks:       stocksUc,
		transactions: transactionsUc,
		sell:         orderUc,
		authUseCase:  authUseCase,
	}
	newStocksController.rg.POST("/auth", newStocksController.UserAuth)
	protectedGroup := newStocksController.rg.Group("/order", tokenMdw.RequireToken())
	protectedGroup.GET("/name", newStocksController.GetById)
	protectedGroup.POST("/buy", newStocksController.BuyStocks)
	protectedGroup.POST("/sell", newStocksController.CreateNewOrderSell)
	return &newStocksController
}

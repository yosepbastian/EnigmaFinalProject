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
	sellAndBuy   usecase.OrderUseCase
	authUseCase  usecase.AuthUseCase
	portoUsecase usecase.PortfoliosUseCase
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

func (pc *StocksController) BuyStocks(ctx *gin.Context) {
	var request models.OrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := pc.sellAndBuy.OrderBuy(request.UserID, request.Email, request.StockName, request.Quantity, request.Price)
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

func (oc *StocksController) CreateNewOrderSell(ctx *gin.Context) {
	var newSell models.Transaction

	if err := ctx.ShouldBindJSON(&newSell); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := oc.sellAndBuy.OrderSell(&newSell)
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

func (oc *StocksController) GetPortfolio(ctx *gin.Context) {

	var userId models.PortoUserID

	err := ctx.ShouldBindJSON(&userId)

	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
	} else {
		newAsset, err := oc.portoUsecase.GetPortoByUserId(userId.UserID)

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

func NewStocksController(routerGroup *gin.RouterGroup, stocksUc usecase.StocksUseCase, orderUc usecase.OrderUseCase, authUseCase usecase.AuthUseCase, tokenMdw middleware.AuthTokenMiddleWare, portoUsecase usecase.PortfoliosUseCase) *StocksController {
	newStocksController := StocksController{
		rg:           routerGroup,
		stocks:       stocksUc,
		sellAndBuy:   orderUc,
		authUseCase:  authUseCase,
		portoUsecase: portoUsecase,
	}
	newStocksController.rg.POST("/auth", newStocksController.UserAuth)

	protectedGroup := newStocksController.rg.Group("/order", tokenMdw.RequireToken())
	protectedGroup.GET("/stocks", newStocksController.GetAll)
	protectedGroup.POST("/buy", newStocksController.BuyStocks)
	protectedGroup.POST("/sell", newStocksController.CreateNewOrderSell)

	protectedGetGroup := newStocksController.rg.Group("/get", tokenMdw.RequireToken())
	protectedGetGroup.GET("/portfolio", newStocksController.GetPortfolio)

	return &newStocksController
}

package delivery

import (
	"kel1-stockbite-projects/config"
	"kel1-stockbite-projects/delivery/controller"
	"kel1-stockbite-projects/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine         *gin.Engine
	useCaseManager manager.UseCaseManager
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		engine:         ginEngine,
		useCaseManager: usecase,
	}
}

func (a *appServer) initHandlers() {
	controller.NewStocksController(a.engine, a.useCaseManager.StocksUseCase(), a.useCaseManager.BuyStocks(), a.useCaseManager.OrderUseCase())

}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}

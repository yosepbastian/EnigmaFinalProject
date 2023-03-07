package delivery

import (
	"kel1-stockbite-projects/config"
	"kel1-stockbite-projects/delivery/controller"
	"kel1-stockbite-projects/delivery/middleware"
	"kel1-stockbite-projects/manager"
	"kel1-stockbite-projects/usecase"
	"kel1-stockbite-projects/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine         *gin.Engine
	authUseCase    usecase.AuthUseCase
	useCaseManager manager.UseCaseManager
	tokenService   authenticator.AccessToken
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	tokenService := authenticator.NewAccessToken(config.TokenConfig)
	use_case := manager.NewUseCaseManager(repo)
	authUserCase := usecase.NewAuthUseCase(tokenService, repo.UsersRepository())

	return &appServer{
		engine:         ginEngine,
		useCaseManager: use_case,
		authUseCase:    authUserCase,
	}
}

func (a *appServer) initHandlers() {
	publicRoute := a.engine.Group("/login")
	tokenMdw := middleware.NewTokenValidator(a.tokenService)
	controller.NewStocksController(publicRoute, a.useCaseManager.StocksUseCase(), a.useCaseManager.OrderUseCase(), a.authUseCase, tokenMdw, a.useCaseManager.PortfoliosUseCase())

}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}

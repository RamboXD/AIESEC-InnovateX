package main

import (
	"log"
	"net/http"

	"github.com/RamboXD/SRS/controllers"
	"github.com/RamboXD/SRS/initializers"
	"github.com/RamboXD/SRS/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	CityController      controllers.CityController
	CityRouteController routes.CityRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
	BoardController      controllers.BoardController
	BoardRouteController routes.BoardRouteController
	CompanyController      controllers.CompanyController
	CompanyRouteController routes.CompanyRouteController

	GameController      controllers.GameController
	GameRouteController routes.GameRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	CityController = controllers.NewCityController(initializers.DB)
	CityRouteController = routes.NewRouteCityController(CityController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)


	BoardController = controllers.NewBoardController(initializers.DB)
	BoardRouteController = routes.NewRouteBoardController(BoardController)

	CompanyController = controllers.NewCompanyController(initializers.DB)
	CompanyRouteController = routes.NewRouteCompanyController(CompanyController)

	GameController = controllers.NewGameController(initializers.DB)
	GameRouteController = routes.NewRouteGameController(GameController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:7000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	CityRouteController.CityRoute(router)
	UserRouteController.UserRoute(router)
	BoardRouteController.BoardRoute(router)
	CompanyRouteController.CompanyRoute(router)
	GameRouteController.GameRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}


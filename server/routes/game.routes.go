package routes

import (
	"github.com/RamboXD/SRS/controllers"
	"github.com/gin-gonic/gin"
)

type GameRouteController struct {
	gameController controllers.GameController
}

func NewRouteGameController(gameController controllers.GameController) GameRouteController {
	return GameRouteController{gameController}
}

func (gc *GameRouteController) GameRoute(rg *gin.RouterGroup) {
	router := rg.Group("game")
	router.POST("/create", gc.gameController.CreateGame)
}

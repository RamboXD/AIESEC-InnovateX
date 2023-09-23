package routes

import (
	"github.com/RamboXD/SRS/controllers"
	"github.com/gin-gonic/gin"
)

type BoardRouteController struct {
	boardController controllers.BoardController
}

func NewRouteBoardController(boardController controllers.BoardController) BoardRouteController {
	return BoardRouteController{boardController}
}

func (cc *BoardRouteController) BoardRoute(rg *gin.RouterGroup) {

	router := rg.Group("board")
	router.POST("/create", cc.boardController.CreateBoard)
}


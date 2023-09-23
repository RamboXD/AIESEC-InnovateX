package routes

import (
	"github.com/RamboXD/SRS/controllers"
	"github.com/gin-gonic/gin"
)

type CityRouteController struct {
	cityController controllers.CityController
}

func NewRouteCityController(cityController controllers.CityController) CityRouteController {
	return CityRouteController{cityController}
}

func (cc *CityRouteController) CityRoute(rg *gin.RouterGroup) {

	router := rg.Group("city")
	router.POST("/create", cc.cityController.CreateCity)
}


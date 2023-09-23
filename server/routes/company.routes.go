package routes

import (
	"github.com/RamboXD/SRS/controllers"
	"github.com/gin-gonic/gin"
)

type CompanyRouteController struct {
	companyController controllers.CompanyController
}

func NewRouteCompanyController(companyController controllers.CompanyController) CompanyRouteController {
	return CompanyRouteController{companyController}
}

func (cc *CompanyRouteController) CompanyRoute(rg *gin.RouterGroup) {

	router := rg.Group("company")
	router.POST("/create", cc.companyController.CreateCompany)
	router.GET("/", cc.companyController.GetAllCompanies)
}


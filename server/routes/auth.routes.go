package routes

import (
	"github.com/RamboXD/SRS/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/create", rc.authController.CreateUser)
	router.POST("/send", rc.authController.SendCode)
	router.POST("/verify", rc.authController.VerifyCode)
}


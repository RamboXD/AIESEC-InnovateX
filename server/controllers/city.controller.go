package controllers

import (
	"net/http"
	"strings"

	"github.com/RamboXD/SRS/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CityController struct {
	DB *gorm.DB
}

func NewCityController(DB *gorm.DB) CityController {
	return CityController{DB}
}

func (cc *CityController) CreateCity(ctx *gin.Context) {
	var payload *models.CityInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// now := time.Now()
	newCity := models.City{
		Name:      payload.Name,
		Code: 	payload.Code,
	}

	result := cc.DB.Create(&newCity)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	cityResponse := &models.CityResponse{
		Name:      payload.Name,
		Code: 	   payload.Code,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": cityResponse}})
}
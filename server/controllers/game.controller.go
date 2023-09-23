package controllers

import (
	"net/http"
	"strings"

	"github.com/RamboXD/SRS/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameController struct {
	DB *gorm.DB
}

func NewGameController(DB *gorm.DB) GameController {
	return GameController{DB}
}

func (gc *GameController) CreateGame(ctx *gin.Context) {
	var payload *models.GameInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newGame := models.Game {
		Name: payload.Name,
		Min_points: payload.Min_points,
	}

	result := gc.DB.Create(&newGame)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	gameResponse := &models.GameResponse{
		ID: newGame.ID,
		Name: newGame.Name,
		Min_points: newGame.Min_points,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"game": gameResponse}})
}
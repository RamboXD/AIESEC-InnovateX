package controllers

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/RamboXD/SRS/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BoardController struct {
	DB *gorm.DB
}

func NewBoardController(DB *gorm.DB) BoardController {
	return BoardController{DB}
}

func (cc *BoardController) CreateBoard(ctx *gin.Context) {
	var payload *models.BoardInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var cityInstance models.City
	if err := cc.DB.Find(&cityInstance, "name = ?", payload.City).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "City not found"})
		return
	}
	// now := time.Now()
	randomNumber := rand.Intn(1024)
	newBoard := models.Board{
		ID:     cityInstance.Code * 1024 + uint(randomNumber),
		City: 	payload.City,
	}

	result := cc.DB.Create(&newBoard)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	boardResponse := &models.BoardResponse{
		ID:        newBoard.ID,
		City: 	   payload.City,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": boardResponse}})
}
package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RamboXD/SRS/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	twilio "github.com/twilio/twilio-go"

	api "github.com/twilio/twilio-go/rest/api/v2010"
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

func (gc *GameController) GetResult(ctx *gin.Context) {
	var payload *models.ResultInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var gameInstance models.Game
	if err := gc.DB.Find(&gameInstance, "ID = ?", payload.GameID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Game not found"})
		return
	}

	var userInstance models.User
	if err := gc.DB.Find(&userInstance, "phone = ?", payload.UserPhone).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	var companyInstance models.Company
	if err := gc.DB.Find(&companyInstance, "Name = ?", payload.CompanyName).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Company not found"})
		return
	}

	if gameInstance.Min_points > payload.Points {
		resultResponse := models.ResultResponse{
			Result: "lose",
		}

		ctx.JSON(http.StatusAccepted, gin.H{"status": "success", "data": gin.H{"game_report": resultResponse}})
		return
	} 

	resultResponse := models.ResultResponse{
		Result: "win",
		PromoCode: companyInstance.PromoCode,
	}
	ctx.JSON(http.StatusAccepted, gin.H{"status": "success", "data": gin.H{"game_rreport": resultResponse}})

	accountSid := "ACf0002507c5ec9de6b536fc8ae3413837"
	authToken := "11372ad0e341644b868e95d32183b19c"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &api.CreateMessageParams{}

	params.SetFrom("whatsapp:+14155238886")
	// params.SetFrom("+13346001315")
	params.SetBody(fmt.Sprintf(`Congrats! You did it! Your promocode from "%s" is %s`, companyInstance.Name, companyInstance.PromoCode))

	// params.SetTo(payload.UserPhone)
	params.SetTo("whatsapp:" + payload.UserPhone)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
	ctx.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "promo_code sent"})
}


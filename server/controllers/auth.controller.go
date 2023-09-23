package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/RamboXD/SRS/models"
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"

	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) CreateUser(ctx *gin.Context) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Phone: 	   payload.Phone,
		Tries: 		5,
		Last_game: now.Add(-24 * time.Hour),
		CreatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	userResponse := &models.UserResponse{
		ID:         newUser.ID,
		Name:       newUser.Name,
		Phone:      newUser.Phone,
		Tries:      int64(newUser.Tries),
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}


func (ac *AuthController) SendCode(ctx *gin.Context) {
	sendCodeRequest := models.SendCodeRequest{}
	err := ctx.ShouldBindJSON(&sendCodeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	accountSid := "AC37bc2d2839e8db8e15ea2bbbaed54ffe"
	authToken := "60eaee9424046597226cd1bf2089faae"
	verifyServiceSid := "VA1bfe2e11fe86054c77c994352c5255d1"
	


	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateVerificationParams{}
	// params.SetCustomFriendlyName("My First Verify Service")
	params.SetTo(sendCodeRequest.Phone)
	params.SetChannel("whatsapp")

	_, err = client.VerifyV2.CreateVerification(verifyServiceSid, params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "message": "code sent"})
}


func (ac *AuthController) VerifyCode(ctx *gin.Context) {
	verifyCodeRequest := models.VerifyCodeRequest{}
	err := ctx.ShouldBindJSON(&verifyCodeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	accountSid := "AC37bc2d2839e8db8e15ea2bbbaed54ffe"
	authToken := "60eaee9424046597226cd1bf2089faae"
	verifyServiceSid := "VA1bfe2e11fe86054c77c994352c5255d1"
	

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(verifyCodeRequest.Phone)
	params.SetCode(verifyCodeRequest.Code)
	resp, err := client.VerifyV2.CreateVerificationCheck(verifyServiceSid, params)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Twillio error"})
		return
	}
	if *resp.Status != "approved" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Verification failed!"})
		return 
	}

	// if verified check if the user is registered
	var userInstance models.User
	// status := "login"
	if err := ac.DB.Find(&userInstance, "phone = ?", verifyCodeRequest.Phone).Error; err != nil {
		ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"phone": verifyCodeRequest.Phone, "name": ""}})
		return
	}

	//if not registered register user

	//generate access token for farmer with expiration time, jwt secret key and user id
	//generate refresh token for farmer with expiration time, jwt secret key and user id

	//staus might be login, register
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"phone": verifyCodeRequest.Phone, "name": userInstance.Name}})

	// ctx.JSON(http.StatusOK, gin.H{"status": "ok", "message": "code verified"})
}


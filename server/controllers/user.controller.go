package controllers

import (
	"net/http"

	"time"

	"github.com/RamboXD/SRS/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

func isToday(t time.Time) bool {
    today := time.Now()
    return t.Year() == today.Year() && t.Month() == today.Month() && t.Day() == today.Day()
}

func (uc *UserController) CheckForAttempts(ctx *gin.Context) {
	var payload *models.CheckInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var userInstance models.User

	if err := uc.DB.Find(&userInstance, "phone = ?", payload.Phone).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Phone is not found"})
		return
	}

	if isToday(userInstance.Last_game) {
		if userInstance.Tries == 0 {
			response := models.CheckResponse{
				IsAllowed: false,
			}

			ctx.JSON(http.StatusAccepted, gin.H{"status": "success", "data": gin.H{"is_allowed": response.IsAllowed}})
		} else {
			response := models.CheckResponse{
				IsAllowed: true,
			}

			tries := userInstance.Tries

			uc.DB.Model(&userInstance).Update("Tries", tries - 1)
			uc.DB.Model(&userInstance).Update("Last_game", time.Now())

			ctx.JSON(http.StatusAccepted, gin.H{"status": "success", "data": gin.H{"is_allowed": response.IsAllowed}})
		}
	} else {
		response := models.CheckResponse {
			IsAllowed: true,
		}

		uc.DB.Model(&userInstance).Update("Tries", 4)
		uc.DB.Model(&userInstance).Update("Last_game", time.Now())

		ctx.JSON(http.StatusAccepted, gin.H{"status": "accepted", "data": gin.H{"is_allowed": response.IsAllowed}})
	}

}


package controllers

import (
	"net/http"

	"github.com/RamboXD/SRS/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompanyController struct {
	DB *gorm.DB
}

func NewCompanyController(DB *gorm.DB) CompanyController {
	return CompanyController{DB}
}



func (cc *CompanyController) CreateCompany(ctx *gin.Context) {
	var payload *models.CompanyInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	
	newCompany := models.Company{
		Name:          payload.Name,
		PromoCode:     payload.PromoCode,
		PromoCodeDesc: payload.PromoCodeDesc,
		PhotoURL:      payload.PhotoURL,
		PrimaryColor:  payload.PrimaryColor,
		SecondaryColor: payload.SecondaryColor,
	}
	result := cc.DB.Create(&newCompany)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	companyResponse := &models.CompanyResponse{
		Name:          payload.Name,
		PromoCode:     payload.PromoCode,
		PromoCodeDesc: payload.PromoCodeDesc,
		PhotoURL:      payload.PhotoURL,
		PrimaryColor:  payload.PrimaryColor,
		SecondaryColor: payload.SecondaryColor,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success",  "data": gin.H{"company": companyResponse}})
}

// Function to get all companies
func (cc *CompanyController) GetAllCompanies(ctx *gin.Context) {
	var companies []models.Company
	var companiesResponse []models.CompanyResponse

	result := cc.DB.Find(&companies)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, models.CompaniesResponse{
			Status:  "error",
			Message: "Something bad happened",
		})
		return
	}

	for _, company := range companies {
		companiesResponse = append(companiesResponse, models.CompanyResponse{
			Name:          company.Name,
			PromoCode:     company.PromoCode,
			PromoCodeDesc: company.PromoCodeDesc,
			PhotoURL:      company.PhotoURL,
			PrimaryColor:  company.PrimaryColor,
			SecondaryColor: company.SecondaryColor,
		})
	}

	ctx.JSON(http.StatusOK, models.CompaniesResponse{
		Status: "success",
		Data:   companiesResponse,
	})
}

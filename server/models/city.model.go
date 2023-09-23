package models

type City struct {
	Name string `gorm:"type:varchar(255);not null"`
	Code uint   `gorm:"type:uint;primary_key"`
}

type CityInput struct {
	Name           string `json:"name" binding:"required"`
	Code           uint `json:"code" binding:"required"`
}

type CityResponse struct {
	Name      string    `json:"name,omitempty"`
	Code 	  uint `json:"code" binding:"required"`
}
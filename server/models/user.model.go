package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"type:uint;primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(255);not null"`
	Tries     int       `gorm:"type:int;not null"`
	Last_game time.Time
	CreatedAt time.Time
}


type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Phone 			string `json:"phone" binding:"required"`
}

type UserResponse struct {
	ID        uint `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Phone 	  string 	`json:"phone" binding:"required"`
	Tries 	  int64		`json:"tries" binding:"required"`
}

type CheckInput struct {
	Phone 	  string 	`json:"phone" binding:"required"`
}

type CheckResponse struct {
	IsAllowed bool      `json:"is_allowed" binding:"required"`
}
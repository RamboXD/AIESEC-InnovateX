package models

type Board struct {
	ID        uint      `gorm:"type:uint;primary_key"`
	City  string    	`gorm:"type:varchar(255);not null"`
}

type BoardInput struct {
	City           string `json:"city" binding:"required"`
}

type BoardResponse struct {
	ID        uint      `gorm:"type:uint;primary_key"`
	City      string    `json:"city,omitempty"`
}




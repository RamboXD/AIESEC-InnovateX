package models

type Game struct {
	ID            uint   `gorm:"type:uint;primary_key"`
	Name          string `gorm:"type:varchar(255);not null"`
	Min_points    int64  `gorm:"type:int;not null"`
}

type GameInput struct {
	Name           string `json:"name" binding:"required"`
	Min_points     int64  `json:"min_points" binding:"required"`
}

type GameResponse struct {
	ID             uint   `json:"id,omitempty"`
	Name           string `json:"name" binding:"required"`
	Min_points     int64  `json:"min_points" binding:"required"`
}


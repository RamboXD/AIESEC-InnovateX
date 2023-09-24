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

type ResultInput struct {
	GameID         uint   `json:"game_id" binding:"required"`
	UserPhone      string `json:"user_phone" binding:"required"`
	CompanyName    string `json:"company_name" binding:"required"`
	Points         int64  `json:"points" binding:"required"`
}

type ResultResponse struct {
	Result         string `json:"result" binding:"required"`
	PromoCode      string `json:"promocode,omitempty"`
}

type SendPromoRequest struct {
	Phone string `json:"phone" binding:"required"`
	Promo string `json:"promo" binding:"required"`
}





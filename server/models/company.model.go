package models

type Company struct {
	Name          string `gorm:"type:varchar(255);not null"`
	PromoCode     string `gorm:"type:varchar(100);"`
	PromoCodeDesc    string `gorm:"type:text;"`  // New field for promo code description
	PhotoURL      string `gorm:"type:varchar(500);"`
	PrimaryColor  string `gorm:"type:varchar(7);"`  // assuming the color is in HEX format
	SecondaryColor string `gorm:"type:varchar(7);"` // assuming the color is in HEX format
}

type CompanyInput struct {
	Name          string `json:"name" binding:"required"`
	PromoCode     string `json:"promo_code,omitempty"`
	PromoCodeDesc string `json:"promo_code_desc,omitempty"`
	PhotoURL      string `json:"photo_url,omitempty"`
	PrimaryColor  string `json:"primary_color,omitempty"`
	SecondaryColor string `json:"secondary_color,omitempty"`
}

// CompanyResponse is used to format the JSON response
type CompanyResponse struct {
	Name          string `json:"name,omitempty"`
	PromoCode     string `json:"promo_code,omitempty"`
	PromoCodeDesc string `json:"promo_code_desc,omitempty"`
	PhotoURL      string `json:"photo_url,omitempty"`
	PrimaryColor  string `json:"primary_color,omitempty"`
	SecondaryColor string `json:"secondary_color,omitempty"`
}
type CompaniesResponse struct {
	Status  string             `json:"status"`
	Data    []CompanyResponse  `json:"data"`
	Message string             `json:"message,omitempty"`
}

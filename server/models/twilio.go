package models

type SendCodeRequest struct {
	Phone       string `validate:"required,min=2,max=100" json:"phone"`
	DeviceToken string `json:"device_token"`
}

type VerifyCodeRequest struct {
	Phone string `validate:"required,min=2,max=100" json:"phone"`
	Code  string `validate:"required,min=4,max=10" json:"code"`
}
package model

type RegisterCustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	NIK         string `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

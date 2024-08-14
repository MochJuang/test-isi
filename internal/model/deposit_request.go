package model

type DepositRequest struct {
	CardNumber string `json:"card_number" validate:"required"`
	Nominal    int64  `json:"nominal" validate:"required"`
}

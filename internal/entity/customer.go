package entity

type Customer struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	NIK         string `json:"nik" gorm:"unique"`
	PhoneNumber string `json:"phone_number" gorm:"unique"`
	Account     Account
}

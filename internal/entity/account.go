package entity

type Account struct {
	ID         uint   `gorm:"primaryKey"`
	CustomerID uint   `gorm:"unique"`
	CardNumber string `json:"card_number" gorm:"unique"`
	Balance    int64  `json:"balance"`
}

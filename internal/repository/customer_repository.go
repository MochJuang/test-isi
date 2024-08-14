package repository

import "test-isi/internal/entity"

type CustomerRepository interface {
	CreateCustomer(customer *entity.Customer) error
	GetCustomerByNIK(nik string) (*entity.Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*entity.Customer, error)
	GetAccountByCardNumber(cardNumber string) (*entity.Account, error)
	UpdateAccount(account *entity.Account) error
}

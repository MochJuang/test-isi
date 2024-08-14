package postgresql

import (
	"gorm.io/gorm"
	"test-isi/internal/entity"
	"test-isi/internal/repository"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) CreateCustomer(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetCustomerByNIK(nik string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where("nik = ?", nik).First(&customer).Error
	return &customer, err
}

func (r *customerRepository) GetCustomerByPhoneNumber(phoneNumber string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where("phone_number = ?", phoneNumber).First(&customer).Error
	return &customer, err
}

func (r *customerRepository) GetAccountByCardNumber(cardNumber string) (*entity.Account, error) {
	var account entity.Account
	err := r.db.Where("card_number = ?", cardNumber).First(&account).Error
	return &account, err
}

func (r *customerRepository) UpdateAccount(account *entity.Account) error {
	return r.db.Save(account).Error
}

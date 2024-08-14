package service

import (
	"errors"
	"fmt"
	"test-isi/internal/entity"
	"test-isi/internal/exception"
	"test-isi/internal/model"
	"test-isi/internal/repository"
	"test-isi/internal/utils"
)

type CustomerService interface {
	RegisterCustomer(req model.RegisterCustomerRequest) (*model.RegisterCustomerResponse, error)
	Deposit(req model.DepositRequest) (*model.DepositResponse, error)
	Withdraw(req model.WithdrawRequest) (*model.WithdrawResponse, error)
	GetBalance(cardNumber string) (*model.GetBalanceResponse, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) RegisterCustomer(req model.RegisterCustomerRequest) (*model.RegisterCustomerResponse, error) {
	if err := utils.Validate(req); err != nil {
		return nil, exception.Validation(err)
	}

	if _, err := s.repo.GetCustomerByNIK(req.NIK); err == nil {
		return nil, exception.Validation(errors.New("NIK already used"))
	}
	if _, err := s.repo.GetCustomerByPhoneNumber(req.PhoneNumber); err == nil {
		return nil, exception.Validation(errors.New("phone Number already used"))
	}

	customer := &entity.Customer{
		Name:        req.Name,
		NIK:         req.NIK,
		PhoneNumber: req.PhoneNumber,
		Account: entity.Account{
			CardNumber: fmt.Sprintf("ACC%v", req.NIK),
			Balance:    0,
		},
	}

	if err := s.repo.CreateCustomer(customer); err != nil {
		return nil, exception.Internal(err)
	}

	return &model.RegisterCustomerResponse{CardNumber: customer.Account.CardNumber}, nil
}

func (s *customerService) Deposit(req model.DepositRequest) (*model.DepositResponse, error) {
	if err := utils.Validate(req); err != nil {
		return nil, exception.Validation(err)
	}

	account, err := s.repo.GetAccountByCardNumber(req.CardNumber)
	if err != nil {
		return nil, exception.NotFound("account not recognized")
	}

	account.Balance += req.Nominal
	if err := s.repo.UpdateAccount(account); err != nil {
		return nil, exception.Internal(err)
	}

	return &model.DepositResponse{Balance: account.Balance}, nil
}

func (s *customerService) Withdraw(req model.WithdrawRequest) (*model.WithdrawResponse, error) {
	if err := utils.Validate(req); err != nil {
		return nil, exception.Validation(err)
	}

	account, err := s.repo.GetAccountByCardNumber(req.CardNumber)
	if err != nil {
		return nil, exception.NotFound("account not recognized")
	}

	if account.Balance < req.Nominal {
		return nil, exception.Validation(errors.New("insufficient balance"))
	}

	account.Balance -= req.Nominal
	if err := s.repo.UpdateAccount(account); err != nil {
		return nil, exception.Internal(err)
	}

	return &model.WithdrawResponse{Balance: account.Balance}, nil
}

func (s *customerService) GetBalance(cardNumber string) (*model.GetBalanceResponse, error) {
	account, err := s.repo.GetAccountByCardNumber(cardNumber)
	if err != nil {
		return nil, exception.NotFound("account not recognized")
	}

	return &model.GetBalanceResponse{Balance: account.Balance}, nil
}

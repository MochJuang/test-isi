package http

import (
	"github.com/gofiber/fiber/v2"
	"log"
	e "test-isi/internal/exception"
	"test-isi/internal/model"
	"test-isi/internal/service"
)

type CustomerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (h *CustomerController) RegisterCustomer(c *fiber.Ctx) error {
	var req model.RegisterCustomerRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request: %v", err)
		return e.Validation(err)
	}

	resp, err := h.service.RegisterCustomer(req)
	if err != nil {
		log.Printf("Error registering customer: %v", err)
		return err
	}

	log.Printf("Customer registered: %v", resp.CardNumber)
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *CustomerController) Deposit(c *fiber.Ctx) error {
	var req model.DepositRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request: %v", err)
		return e.Validation(err)
	}

	resp, err := h.service.Deposit(req)
	if err != nil {
		log.Printf("Error depositing: %v", err)
		return err
	}

	log.Printf("Deposit successful: %v", resp.Balance)
	return c.JSON(resp)
}

func (h *CustomerController) Withdraw(c *fiber.Ctx) error {
	var req model.WithdrawRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request: %v", err)
		return e.Validation(err)
	}

	resp, err := h.service.Withdraw(req)
	if err != nil {
		log.Printf("Error withdrawing: %v", err)
		return err
	}

	log.Printf("Withdraw successful: %v", resp.Balance)
	return c.JSON(resp)
}

func (h *CustomerController) GetBalance(c *fiber.Ctx) error {
	cardNumber := c.Params("account_no")

	resp, err := h.service.GetBalance(cardNumber)
	if err != nil {
		log.Printf("Error getting balance: %v", err)
		return err
	}

	log.Printf("Balance retrieved: %v", resp.Balance)
	return c.JSON(resp)
}

package route

import (
	"github.com/gofiber/fiber/v2"
	"test-isi/internal/config"
	httpdelivery "test-isi/internal/delivery/http"
	middleware "test-isi/internal/delivery/http/midlleware"
	postgresqlrepository "test-isi/internal/repository/postgresql"
	"test-isi/internal/service"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {
	// Initialize http
	app.Use(middleware.ErrorHandlerMiddleware)

	// Public routes
	customerRepo := postgresqlrepository.NewCustomerRepository(cfg.DB)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := httpdelivery.NewCustomerController(customerService)

	app.Post("/daftar", customerHandler.RegisterCustomer)
	app.Post("/tabung", customerHandler.Deposit)
	app.Post("/tarik", customerHandler.Withdraw)
	app.Get("/saldo/:account_no", customerHandler.GetBalance)
}

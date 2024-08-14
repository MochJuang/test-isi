package middleware

import (
	e "test-isi/internal/exception"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandlerMiddleware(c *fiber.Ctx) error {

	err := c.Next()

	if err != nil {
		return e.HandleHttpErrorFiber(c, err)
	}

	return nil
}

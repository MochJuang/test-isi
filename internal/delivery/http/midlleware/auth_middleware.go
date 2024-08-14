package middleware

import (
	"errors"
	"fmt"
	"strings"
	"test-isi/internal/config"
	e "test-isi/internal/exception"
	"test-isi/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("from auth middleware")
		token := c.Get("Authorization")
		if token == "" {
			err := e.Unauthorized(errors.New("missing token"))
			return e.HandleHttpErrorFiber(c, err)
		}
		token = strings.Replace(token, "Bearer ", "", 1)

		claims, err := utils.ParseToken(token, cfg.JWTSecret)
		if err != nil {
			err = e.Unauthorized(err)
			return e.HandleHttpErrorFiber(c, err)
		}
		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}

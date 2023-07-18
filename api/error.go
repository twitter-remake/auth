package api

import (
	"errors"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Error handles errors returned from handlers
func Error(c *fiber.Ctx, err error) error {
	if err.Error() == "Method Not Allowed" {
		return c.Status(http.StatusMethodNotAllowed).JSON(fiber.Map{
			"message": "Method not allowed",
		})
	}

	log.Err(err).Send()

	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if os.Getenv("ENVIRONMENT") == "dev" {
		if code == fiber.StatusInternalServerError {
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   http.StatusText(code),
		"message": e.Message,
	})
}

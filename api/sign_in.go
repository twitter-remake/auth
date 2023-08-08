package api

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/twitter-remake/auth/backend"
)

type RegisterInput struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	BirthDate  string `json:"birth_date"`
}

func (s *Server) SignIn(c *fiber.Ctx) error {
	var input RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := input.Validate(); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	authorization := c.Get("Authorization")
	if len(strings.Split(authorization, " ")) < 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized")
	}

	idToken := strings.Split(authorization, " ")[1]

	// No need to validate the birth date here since it's already validated
	// in the RegisterInput.Validate() method.
	birthDate, _ := time.Parse("2006-01-02", input.BirthDate)

	err := s.backend.SignIn(c.Context(), backend.RegisterInput{
		IDToken:    idToken,
		Name:       input.Name,
		ScreenName: input.ScreenName,
		BirthDate:  birthDate,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Signed In.",
	})
}

func (i RegisterInput) Validate() error {
	if i.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Name is required")
	}

	if i.ScreenName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Screen name is required")
	}

	_, err := time.Parse("2006-01-02", i.BirthDate)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid birth date")
	}

	return nil
}

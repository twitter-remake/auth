package api

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/twitter-remake/auth/backend"
)

type RegisterInput struct {
	Name       string    `json:"name"`
	ScreenName string    `json:"screen_name"`
	Location   string    `json:"location"`
	BirthDate  time.Time `json:"birth_date"`
}

func (s *Server) SignIn(c *fiber.Ctx) error {
	var input RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	authorization := c.Get("Authorization")
	if len(strings.Split(authorization, " ")) < 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized")
	}

	idToken := strings.Split(authorization, " ")[1]

	err := s.backend.SignIn(c.Context(), backend.RegisterInput{
		IDToken:    idToken,
		Name:       input.Name,
		ScreenName: input.ScreenName,
		Location:   input.Location,
		BirthDate:  input.BirthDate,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

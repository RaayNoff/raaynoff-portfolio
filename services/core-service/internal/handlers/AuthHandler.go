package handlers

import (
	dtos "core-service/internal/dtos/auth"
	"core-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login godoc
// @Summary Login
// @Description Login with credentials
// @Tags auth
// @Accept  json
// @Produce  json
// @Param AuthLogin body dtos.AuthLoginDto true "User data"
// @Success 201 {object} dtos.AuthLoginSchema
// @Router /auth/login [post]
func (a *AuthHandler) Login(c *fiber.Ctx) error {
	var authLoginDto dtos.AuthLoginDto

	if err := c.BodyParser(&authLoginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	a.service.Login(&authLoginDto)
}

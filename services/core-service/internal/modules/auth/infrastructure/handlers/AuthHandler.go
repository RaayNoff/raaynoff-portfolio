package authHandlers

import (
	authDtos "core-service/internal/modules/auth/domain/dtos"
	authServices "core-service/internal/modules/auth/domain/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *authServices.AuthService
}

func NewAuthHandler(svc *authServices.AuthService) *AuthHandler {
	return &AuthHandler{service: svc}
}

// Login godoc
// @Summary Login
// @Description Login with credentials
// @Tags auth
// @Accept  json
// @Produce  json
// @Param AuthLogin body authDtos.AuthLoginDto true "User data"
// @Success 201 {object} authSchemas.AuthLoginSchema
// @Router /auth/login [post]
func (a *AuthHandler) Login(c *fiber.Ctx) error {
	var authLoginDto authDtos.AuthLoginDto

	if err := c.BodyParser(&authLoginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	a.service.Login(&authLoginDto)

	return nil
}

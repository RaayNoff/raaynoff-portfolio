package routes

import (
	"core-service/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router, handler *handlers.AuthHandler) {
	route := router.Group("/auth")
	route.Post("/login", handler.Login)
}

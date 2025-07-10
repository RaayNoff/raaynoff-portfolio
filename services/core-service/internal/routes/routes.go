package routes

import (
	authHandlers "core-service/internal/modules/auth/infrastructure/handlers"
	blogHandlers "core-service/internal/modules/blog/infrastructure/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	blogHandler *blogHandlers.BlogHandler,
	authHandler *authHandlers.AuthHandler,
) {

	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("health route works!")
	})

	// Блог
	blogGroup := api.Group("/blog")
	app.Get("/", blogHandler.GetAll)
	blogGroup.Post("/", blogHandler.Add)

	// Аутентификация
	authGroup := api.Group("/auth")
	authGroup.Post("/login", authHandler.Login)
}

package routes

import (
	"core-service/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupBlogRoutes(router fiber.Router, handler *handlers.BlogHandler) {
	route := router.Group("/blog")
	route.Get("/", handler.GetAll)
	route.Post("/", handler.Add)
}

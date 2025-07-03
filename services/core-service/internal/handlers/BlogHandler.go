package handlers

import (
	dtos "core-service/internal/dtos/blog"
	"core-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type BlogHandler struct {
	service services.BlogService
}

func NewBlogHandler(service services.BlogService) *BlogHandler {
	return &BlogHandler{service}
}

// GetAll godoc
// @Summary Get all blog posts
// @Description Get all blog posts from database
// @Tags blog
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Blog
// @Router /blog [get]
func (h *BlogHandler) GetAll(c *fiber.Ctx) error {
	posts, err := h.service.GetAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(posts)
}

// Add godoc
// @Summary Create blog post
// @Description Creates a new blog post
// @Tags blog
// @Accept  json
// @Produce  json
// @Param blog body dtos.CreateBlogDto true "Blog post data"
// @Success 201 {object} models.Blog
// @Router /blog [post]
func (h *BlogHandler) Add(c *fiber.Ctx) error {
	var input dtos.CreateBlogDto

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

package blogHandlers

import (
	blogDtos "core-service/internal/modules/blog/domain/dtos"
	blogServices "core-service/internal/modules/blog/domain/services"
	"github.com/gofiber/fiber/v2"
)

type BlogHandler struct {
	service *blogServices.BlogService
}

func NewBlogHandler(svc *blogServices.BlogService) *BlogHandler {
	return &BlogHandler{service: svc}
}

// GetAll godoc
// @Summary Get all blog posts
// @Description Get all blog posts from database
// @Tags blog
// @Accept  json
// @Produce  json
// @Success 200 {array} blogModels.Blog
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
// @Param blog body blogDtos.BlogSaveDto true "Blog post data"
// @Success 201 {object} blogModels.Blog
// @Router /blog [post]
func (h *BlogHandler) Add(c *fiber.Ctx) error {
	var input blogDtos.BlogSaveDto

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

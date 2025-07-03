package containers

import (
	"core-service/internal/handlers"
	"core-service/internal/repositories"
	"core-service/internal/services"
	"gorm.io/gorm"
)

type BlogContainer struct {
	BlogHandler *handlers.BlogHandler
}

func NewBlogContainer(db *gorm.DB) *BlogContainer {
	blogRepo := repositories.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepo)
	blogHandler := handlers.NewBlogHandler(blogService)

	return &BlogContainer{
		BlogHandler: blogHandler,
	}
}

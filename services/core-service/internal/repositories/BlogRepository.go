package repositories

import (
	"core-service/internal/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetAll() ([]models.Blog, error)
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}

func (repository *blogRepository) GetAll() ([]models.Blog, error) {
	var posts []models.Blog

	err := repository.db.Find(&posts).Error

	return posts, err
}

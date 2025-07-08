package blogRepositories

import (
	blogModels "core-service/internal/modules/blog/domain/models"
	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (repository *BlogRepository) GetAll() ([]blogModels.Blog, error) {
	var posts []blogModels.Blog

	err := repository.db.Find(&posts).Error

	return posts, err
}

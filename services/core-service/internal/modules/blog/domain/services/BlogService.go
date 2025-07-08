package blogServices

import (
	blogModels "core-service/internal/modules/blog/domain/models"
	blogRepositories "core-service/internal/modules/blog/infrastructure/repositories"
)

type BlogService struct {
	repo *blogRepositories.BlogRepository
}

func NewBlogService(repo *blogRepositories.BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (service *BlogService) GetAll() ([]blogModels.Blog, error) {
	return service.repo.GetAll()
}

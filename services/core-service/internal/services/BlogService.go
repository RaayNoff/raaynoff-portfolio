package services

import (
	"core-service/internal/models"
	"core-service/internal/repositories"
)

type BlogService interface {
	GetAll() ([]models.Blog, error)
}

type blogService struct {
	repo repositories.BlogRepository
}

func NewBlogService(repo repositories.BlogRepository) BlogService {
	return &blogService{repo}
}

func (service *blogService) GetAll() ([]models.Blog, error) {
	return service.repo.GetAll()
}

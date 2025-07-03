package services

import (
	"core-service/internal/models"
	"core-service/internal/repositories"
)

type UserService interface {
	FindByLogin(login string) (models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (service *userService) FindByLogin(login string) (models.User, error) {
	return service.repo.FindByLogin(login)
}

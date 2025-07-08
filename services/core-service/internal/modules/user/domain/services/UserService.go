package userServices

import (
	userModels "core-service/internal/modules/user/domain/models"
	userRepositories "core-service/internal/modules/user/infrastructure/repositories"
)

type UserService struct {
	repo *userRepositories.UserRepository
}

func NewUserService(repo *userRepositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) FindByLogin(login string) (userModels.User, error) {
	return service.repo.FindByLogin(login)
}

package userRepositories

import (
	userModels "core-service/internal/modules/user/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) FindByLogin(login string) (userModels.User, error) {
	var user userModels.User

	err := repository.db.Where("login = ?", login).Find(&user).Error

	return user, err
}

package repositories

import (
	"core-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByLogin(login string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) FindByLogin(login string) (models.User, error) {
	var user models.User

	err := repository.db.Where("login = ?", login).Find(&user).Error

	return user, err
}

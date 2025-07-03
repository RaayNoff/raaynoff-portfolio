package containers

import (
	"core-service/internal/repositories"
	"core-service/internal/services"
	"gorm.io/gorm"
)

type UserContainer struct {
	UserService    services.UserService
	UserRepository *repositories.UserRepository
}

func NewUserContainer(db *gorm.DB) *UserContainer {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	return &UserContainer{
		UserService: userService,
	}
}

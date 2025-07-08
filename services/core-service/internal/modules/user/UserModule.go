package userModule

import (
	userServices "core-service/internal/modules/user/domain/services"
	userRepositories "core-service/internal/modules/user/infrastructure/repositories"
	"go.uber.org/dig"
)

// UserModule – DI-модуль для user.
type UserModule struct {
	Service *userServices.UserService
}

// NewUserModule создаёт и возвращает UserModule с зависимостями.
func NewUserModule(
	svc *userServices.UserService,
) *UserModule {
	return &UserModule{
		Service: svc,
	}
}

// Register регестрирует зависимости модуля в DI-контейнере.
func Register(container *dig.Container) error {
	// Репозиторий
	if err := container.Provide(userRepositories.NewUserRepository); err != nil {
		return err
	}

	// Сервис
	if err := container.Provide(userServices.NewUserService); err != nil {
		return err
	}

	// Сам модуль
	return container.Provide(NewUserModule)
}

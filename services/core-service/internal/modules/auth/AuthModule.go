package authModule

import (
	authServices "core-service/internal/modules/auth/domain/services"
	authHandlers "core-service/internal/modules/auth/infrastructure/handlers"
	"go.uber.org/dig"
)

// AuthModule – DI-модуль для auth.
type AuthModule struct {
	Handler *authHandlers.AuthHandler
	Service *authServices.AuthService
}

// NewAuthModule создаёт и возвращает AuthModule с зависимостями.
func NewAuthModule(
	svc *authServices.AuthService,
	handler *authHandlers.AuthHandler,
) *AuthModule {
	return &AuthModule{
		Handler: handler,
		Service: svc,
	}
}

// Register регестрирует зависимости модуля в DI-контейнере.
func Register(container *dig.Container) error {
	// Сервис
	if err := container.Provide(authServices.NewAuthService); err != nil {
		return err
	}

	// Хендлер
	if err := container.Provide(authHandlers.NewAuthHandler); err != nil {
		return err
	}

	// Сам модуль
	return container.Provide(NewAuthModule)
}

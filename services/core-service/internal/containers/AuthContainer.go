package containers

import (
	"core-service/config"
	"core-service/internal/handlers"
	"core-service/internal/services"
)

type AuthContainer struct {
	AuthService *services.AuthService
	AuthHandler *handlers.AuthHandler
}

func NewAuthContainer(cfg *config.Config) *AuthContainer {

	authService := services.NewAuthService(cfg.JwtSecret)
	authHandler := handlers.NewAuthHandler(*authService)

	return &AuthContainer{
		AuthService: authService,
		AuthHandler: authHandler,
	}
}

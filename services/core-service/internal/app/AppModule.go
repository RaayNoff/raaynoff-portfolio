package app

import (
	"core-service/config"
	authModule "core-service/internal/modules/auth"
	blogModule "core-service/internal/modules/blog"
	userModule "core-service/internal/modules/user"
	"fmt"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

// AppModule – главный DI-контейнер приложения.
type AppModule struct {
	Container *dig.Container
}

// NewAppModule создаёт и настраивает DI-контейнер.
func NewAppModule(db *gorm.DB, cfg *config.Config) *AppModule {
	container := dig.New()

	container.Provide(func() *gorm.DB {
		return db
	})

	container.Provide(func() *config.Config {
		return cfg
	})

	registerFns := []func(*dig.Container) error{
		blogModule.Register,
		authModule.Register,
		userModule.Register,
	}

	for _, register := range registerFns {
		if err := register(container); err != nil {
			fmt.Printf("Failed to register module: %v", err)
		}
	}

	return &AppModule{Container: container}
}

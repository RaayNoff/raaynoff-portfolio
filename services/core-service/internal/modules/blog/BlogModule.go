package blogModule

import (
	blogServices "core-service/internal/modules/blog/domain/services"
	blogHandlers "core-service/internal/modules/blog/infrastructure/handlers"
	blogRepositories "core-service/internal/modules/blog/infrastructure/repositories"
	"go.uber.org/dig"
)

// BlogModule – DI-модуль для блога.
type BlogModule struct {
	Handler    *blogHandlers.BlogHandler
	Service    *blogServices.BlogService
	Repository *blogRepositories.BlogRepository
}

// NewBlogModule создаёт и возвращает BlogModule с зависимостями.
func NewBlogModule(
	repo *blogRepositories.BlogRepository,
	svc *blogServices.BlogService,
	handler *blogHandlers.BlogHandler,
) *BlogModule {
	return &BlogModule{
		Handler:    handler,
		Service:    svc,
		Repository: repo,
	}
}

// Register регестрирует зависимости модуля в DI-контейнере.
func Register(container *dig.Container) error {
	// Репозиторий
	if err := container.Provide(blogRepositories.NewBlogRepository); err != nil {
		return err
	}

	// Сервис
	if err := container.Provide(blogServices.NewBlogService); err != nil {
		return err
	}

	// Хендлер
	if err := container.Provide(blogHandlers.NewBlogHandler); err != nil {
		return err
	}

	// Сам модуль
	return container.Provide(NewBlogModule)
}

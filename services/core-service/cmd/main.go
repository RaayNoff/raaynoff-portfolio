package main

import (
	"core-service/config"
	_ "core-service/docs"
	"core-service/internal/containers"
	"core-service/internal/handlers"
	"core-service/internal/models"
	"core-service/internal/repositories"
	"core-service/internal/routes"
	"core-service/internal/services"
	"core-service/pkg/database"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

// @title Core Service API
// @version 1.0
// @description Core Service API
// @host localhost:3000
// @BasePath /api/v1
func main() {
	cfg := config.Load()
	db := database.Connect(cfg.PostgresDSN)

	app := fiber.New(fiber.Config{
		AppName: "http://localhost:3000/api/docs",
	})

	configureCors(app)
	initSwagger(app)
	api := setGlobalPrefix(app)
	migrateDatabase(db)

	blogContainer := containers.NewBlogContainer(db)
	authContainer := containers.NewAuthContainer(cfg)
	containers.NewUserContainer(db)

	routes.SetupBlogRoutes(api, blogContainer.BlogHandler)
	routes.SetupAuthRoutes(api, authContainer.AuthHandler)
	startServer(app, cfg)
}

func buildContainer(cfg *config.Config) *dig.Container {
	container := dig.New()

	// Provide конфиг и базу
	container.Provide(func() *config.Config {
		return cfg
	})
	container.Provide(func(cfg *config.Config) *gorm.DB {
		return database.Connect(cfg.PostgresDSN)
	})

	// Provide репозитории
	container.Provide(repositories.NewUserRepository)
	container.Provide(repositories.NewBlogRepository)

	// Provide сервисы
	container.Provide(services.NewUserService)
	container.Provide(services.NewBlogService)
	container.Provide(services.NewAuthService)

	// Provide хендлеры
	container.Provide(handlers.NewAuthHandler)
	container.Provide(handlers.NewBlogHandler)

	return container
}

func migrateDatabase(db *gorm.DB) {
	blogMigrate := db.AutoMigrate(&models.Blog{})
	userMigrate := db.AutoMigrate(&models.User{})

	if blogMigrate != nil {
		return
	}
	if userMigrate != nil {
		return
	}

}

func startServer(app *fiber.App, config *config.Config) {
	listenError := app.Listen(":" + config.Port)

	if listenError != nil {
		return
	}
}

func initSwagger(app *fiber.App) {
	cfg := config.GetSwaggerConfig()

	app.Use(swagger.New(*cfg))
}

func configureCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
}

func setGlobalPrefix(app *fiber.App) fiber.Router {
	return app.Group("/api/v1")
}

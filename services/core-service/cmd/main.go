package main

import (
	"core-service/config"
	"core-service/internal/app"
	authModule "core-service/internal/modules/auth"
	blogModule "core-service/internal/modules/blog"
	blogModels "core-service/internal/modules/blog/domain/models"
	userModels "core-service/internal/modules/user/domain/models"
	"core-service/internal/routes"
	"core-service/pkg/database"
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	appInstance := fiber.New(fiber.Config{
		AppName: "http://localhost:3000/api/docs",
	})

	configureCors(appInstance)
	initSwagger(appInstance)
	migrateDatabase(db)

	// 1. Инициализируем DI-контейнер
	appModule := app.NewAppModule(db, cfg)

	// 3. Настраиваем маршруты через DI
	err := appModule.Container.Invoke(func(
		blogModule *blogModule.BlogModule,
		authModule *authModule.AuthModule,
	) {
		routes.SetupRoutes(appInstance, blogModule.Handler, authModule.Handler)
	})

	if err != nil {
		fmt.Printf("DI error: %v", err)
	}

	startServer(appInstance, cfg)
}

func migrateDatabase(db *gorm.DB) {
	blogMigrate := db.AutoMigrate(&blogModels.Blog{})
	userMigrate := db.AutoMigrate(&userModels.User{})

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

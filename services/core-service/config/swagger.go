package config

import "github.com/gofiber/contrib/swagger"

func GetSwaggerConfig() *swagger.Config {
	return &swagger.Config{
		FilePath: "./docs/swagger.json",
		Path:     "api/docs",
	}
}

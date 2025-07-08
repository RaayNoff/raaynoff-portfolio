# Примерная структура сервиса

```
|   .env
|   .env.example
|   go.mod
|   go.sum
|   README.md|       
+---cmd
|       main.go
|       
+---config
|       config.go
|       swagger.go
|       
+---docs
|       docs.go
|       swagger.json
|       swagger.yaml
|       
+---internal
|   +---app
|   +---middleware
|   +---modules
|   |   +---auth
|   |   |   |   Module.go
|   |   |   +---domain
|   |   |   |   +---dtos
|   |   |   |   \---services
|   |   |   |
|   |   |   \---infrastructure
|   |   |       \---handlers
|   |   |       \---schemas
|   |   |       \---repositories
|   \---routes
|           routes.go
|
\---pkg
    \---database
            postgres.go

```

## Генерация swagger документации
```bash
  swag init -g cmd\main.go
```

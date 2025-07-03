# Примерная структура сервиса

```
/your-project
├── cmd/
│   └── main.go              # Точка входа
├── config/
│   └── config.go            # Загрузка и парсинг конфигурации
├── internal/
│   ├── app/                 # Инициализация приложения (Fiber, middlewares)
│   ├── logger/              # Логгер (zap/slog/logrus)
│   ├── database/            # Подключение к PostgreSQL
│   ├── models/              # Структуры моделей
│   ├── repositories/        # Доступ к данным (ORM/sqlx/gorm/pgx)
│   ├── services/            # Бизнес-логика
│   ├── handlers/            # HTTP-хендлеры
│   └── routes/              # Роутинг
├── migrations/              # SQL-миграции
├── .env                     # Конфигурация
├── go.mod
└── go.sum

```

## Генерация swagger документации
```bash
  swag init -g cmd\main.go
```

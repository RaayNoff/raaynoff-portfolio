# raaynoff-portfolio

## Примерная структура проекта:
```
raaynoff/
├── services/
│   ├── core-service/       # Тот самый Go-сервис (CRUD блога, резюме, пользователей и т.д.)
│   ├── photo-service/      # NestJS (загрузка в S3)
│   └── notification/       # Java Spring (Kafka → уведомления)
│
├── deployments/            # Общие настройки
│   ├── docker-compose.yml  # Запуск ВСЕХ сервисов
│   └── kafka/             
│
├── libs/                   # Общий код (если нужен)
│   ├── kafka-events/       # Схемы сообщений для Kafka (DTO)
│   └── s3-client/          # Обёртка для работы с S3
│
└── README.md
```

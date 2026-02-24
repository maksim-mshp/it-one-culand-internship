# Сервис для хранения данных о стажировках

## Запуск в Docker

```bash
docker compose up
```

## Локальный запуск

**Установка зависимостей:**
```bash
task install-deps
```

**Запуск проекта:**
```bash
task run
```

## Миграции

```bash
task migrate -- -help
```

## Генерация OpenAPI
```bash
task openapi
```
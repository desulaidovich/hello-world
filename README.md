# hello-world

Минималистичный скелет HTTP-сервиса на Go. Служит отправной точкой для новых сервисов: DI через `fx`, структурированные логи, конфигурация из YAML + ENV, единый формат ответов.

## Быстрый старт

```bash
go run ./cmd/server --config config.yaml
# или
go run ./cmd/server -c /path/to/config.yaml
```

Сервер поднимается на порту из конфига (по умолчанию `:8080`).

```bash
curl http://localhost:8080/health
# {"status":"OK"}
```

## Конфигурация

Конфиг читается из YAML-файла, потом поверх него накладываются переменные окружения.

**`config.yaml`**

```yaml
development: true

http:
  host: ""
  port: "8080"
  read_header_timeout: 5s
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 60s
```

**Переменные окружения**

Разделитель уровней вложенности — двойное подчёркивание (`__`). Одиночное подчёркивание допустимо в именах полей.

| Переменная                       | Поле конфига                  |
|----------------------------------|-------------------------------|
| `HTTP__PORT`                     | `http.port`                   |
| `HTTP__READ_HEADER_TIMEOUT`      | `http.read_header_timeout`    |
| `HTTP__READ_TIMEOUT`             | `http.read_timeout`           |
| `HTTP__WRITE_TIMEOUT`            | `http.write_timeout`          |
| `HTTP__IDLE_TIMEOUT`             | `http.idle_timeout`           |

ENV имеет приоритет над YAML.

## Структура проекта

```
.
├── cmd/server/         # точка входа, сборка fx-приложения
├── config/             # типы конфига + провайдер (koanf)
└── internal/
    ├── dto/            # общие типы запросов/ответов
    ├── handler/        # HTTP-обработчики
    ├── middleware/     # middleware: request_id, logging, chain
    ├── module/         # fx-модули (Config, Logger, Handler, Server)
    ├── response/       # хелперы записи JSON-ответов
    └── server/         # обёртка над http.Server
```

## Формат ответов

Все ответы — JSON. Структура единая для успехов и ошибок.

**Успех**

```json
{
  "success": true,
  "data": { ... }
}
```

**Ошибка**

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "invalid request",
    "details": [
      { "field": "email", "message": "required" }
    ]
  }
}
```

`details` присутствует только когда нужен.

## Зависимости

| Библиотека | Назначение |
|------------|------------|
| `go.uber.org/fx` | dependency injection |
| `go.uber.org/zap` | структурированные логи |
| `github.com/knadh/koanf` | конфигурация (YAML + ENV) |
| `github.com/go-ozzo/ozzo-validation` | валидация конфига |
| `github.com/google/uuid` | генерация request ID |
| `github.com/urfave/cli` | CLI-флаги (`--config`) |

# USAGE

1. 
```shell
make build-img
```
2.
```shell
make up
```
3. 
```shell
make goose-up # Для поднятия миграций
```

Дополнительно:
```shell
make goose-down # Для отката последней миграции

# Создает миграцию в папке ./internal/repository/migrations/
make goose-create MIGRATION_NAME=name 
```

### Allowed methods and paths
```shell
GET /airticket # Получить список всех билетов
PATCH /airticket/:order_id # Получить информацию о билете
DELETE /airticket/:order_id # Удалить информацию о билете

# Операции над документами пассажира
GET /:passenger_id/documents
PATCH /documents/:document_id
DELETE /documents/:document_id

# Операции над информацией о пассажирах
GET /airticket/:order_id/passengers
PATCH /passengers/:passenger_id
DELETE /passengers/:passenger_id


GET airticket/:order_id/info    # Получить исчерпывающую информацию о билете
GET /passengers/:passenger_id/report    # Получить
```
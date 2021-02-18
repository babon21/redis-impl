# Приложение имплементация in-memory Redis кеша

<!-- ToC start -->
# Обзор

1. [Запуск приложения](#Запуск-приложения)
1. [Юнит-тесты](#Юнит-тесты)
1. [Реализация](#Реализация)
1. [Усложнения](#Усложнения)
1. [Используемые библиотеки](#Используемые-библиотеки)
1. [API клиента](#API-клиента)
1. [API сервера](#API-сервера)
<!-- ToC end -->

# Запуск приложения
```
docker-compose up
``` 

# Юнит-тесты
Запуск юнит-тестов:
```
make test
```

# Реализация
- Язык программирования Go
- HTTP JSON API
# Усложнения
- Реализовать операторы: HGET, HSET, LGET, LSET
- Периодическое удаление истекших ключей в отдельной горутине
- Следование принципам подхода "Чистая архитектура"
- Покрытие кода Юнит тестами
- Докеризация приложения и запуск с помощью `docker-compose`
- Многоэтапная сборка Docker-образа сервиса
- Использование переменных окружения для конфигурации приложения (One of The Twelve-Factor App)
- Использование линтера `golangci-lint`
- Логирование HTTP запросов с помощью middleware
# Используемые библиотеки
- Веб-фреймворк `echo`
- Конфигурирование приложения - библиотека `viper`
- Логирование с помощью `zerolog`

# API клиента
Клиент служит своего рода gateway к серверу.
### GET оператор, POST /cache/string

Возвращает в случае успеха Status 200 и JSON.

Запрос:
```
curl --request GET 'localhost:8081/cache/string/key1'
```
Ответ:
```
{
  "value": "value1"
}
```
### SET оператор, PUT /cache/string
Возвращает в случае успеха Status 201

Запрос:
```
curl --request PUT 'localhost:8081/cache/string' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "key1",
    "value": "value1"
}'
```

### DEL оператор, DELETE /cache/keys
Ответ в случае успеха Status 204

Запрос:
```
curl --request DELETE 'localhost:8081/cache/keys/key1'
```
### KEYS оператор, GET /cache/keys
Возвращает в случае успеха Status 200 и JSON

Запрос:
```
curl --request GET 'localhost:8081/cache/keys' \
--header 'Content-Type: application/json' \
--data-raw '{
    "pattern": ".*name.*"
}'
```
Ответ:
```
{
    "keys": [
        "firstname",
        "lastname"
    ]
}
```
### HSET оператор, PUT /cache/map
Возвращает в случае успеха Status 200 и JSON (кол-во полей, которые были добавлены).

Запрос:
```
curl --request PUT 'localhost:8081/cache/map' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "hkey",
    "pairs": [
        {
            "field": "field1",
            "value": "value1"
        },
        {
            "field": "field2",
            "value": "value2"
        }
    ]
}'
```
Ответ:
```
{
    "count": 2
}
```
### HGET оператор, GET /cache/map
Возвращает в случае успеха Status 200 и JSON.

Запрос:
```
curl --request GET 'localhost:8081/cache/map' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "hkey",
    "field": "field2"
}'
```
Ответ:
```
{
    "value": "value2"
}
```
### LPUSH оператор, POST /cache/list
Возвращает в случае успеха Status 200 и JSON (размер списка после добавления).

Запрос:
```
curl --request POST 'localhost:8081/cache/list' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "lkey",
    "values": [
        "val6",
        "val5"
    ]
}'
```
Ответ:
```
{
    "size": 2
}
```
### LGET оператор, GET /cache/list
Возвращает в случае успеха Status 200 и JSON.

Запрос:
```
curl --request GET 'localhost:8081/cache/list' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "lkey",
    "index": 0
}'
```
Ответ:
```
{
    "value": "val5"
}
```
### LSET оператор, PATCH /cache/list
Возвращает в случае успеха Status 204.

Запрос:
```
curl --request PATCH 'localhost:8081/cache/list' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "lkey",
    "value": "new_val",
    "index": 1
}'
```
### EXPIRE оператор (установка TTL), PATCH /cache/keys/expire
Возвращает в случае успеха Status 204. TTL в запросе указывается в секундах.

Запрос:
```
curl --request PATCH 'localhost:8081/cache/keys/expire' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "lkey",
    "ttl": 10
}'
```

# API сервера
API сервера совпадает с API клиента, для выполнения запросов необходимо изменить только порт. 
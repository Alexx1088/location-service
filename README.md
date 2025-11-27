# Location Service

Сервис для управления городами и улицами.  
Реализован на **Go**, используется **PostgreSQL** и **golang-migrate** для миграций.

##  Развёртывание проекта

### 1. Скачать и установить Go с официального сайта, в соответствии со своей операционной системой:
https://go.dev/dl

#### проверить, что Go работает корректно, ввести команду в терминале:
go version

#### пример вывода:
go version go1.24.1 linux/amd64

### 2. После клонирования репозитория, создать файл с переменными окружения
#### запустить команду в корне проекта:
cp .env.example .env

## 3. Поднять PostgreSQL (через Docker)
###  cборка и запуск контейнеров:

```bash
docker compose build
docker compose up -d
```
### проверить, что контейнеры поднялись: 
docker ps

## 4. Запустить приложение 
### выполнить команду:
go run cmd/location-service/main.go

### после запуска приложение будет доступно по адресу:
http://localhost:8080

## TODO 

1) create the ```Crossing``` entity:
### Crossing
```
- id
- crossroad_id
- event_time
```
2) implement the ```create``` and ```getAll``` REST methods. The ```getAll``` method should include pagination, 
with ```from``` and ```to``` query params for filtering by date and a ```crossroad``` query param for filtering 
by ```crossroad_id```

## Эндпойнты

### Cities:
#### получить город: GET http://localhost:8080/cities/{id}
#### получить список городов: GET http://localhost:8080/cities
#### создать новый город: POST http://localhost:8080/cities
пример запроса в body:
{
"name": "Taraz"
}
#### обновить город: PUT http://localhost:8080/cities/{id}
пример запроса в body:
{
"name": "Taraz"
}
#### удалить город: DELETE http://localhost:8080/cities/{id}

### Streets:
#### получить улицу: GET http://localhost:8080/streets/{id}
#### получить список улиц: GET http://localhost:8080/streets
#### создать новую улицу: POST http://localhost:8080/streets
пример запроса в body:
{
"name": "Gogolya"
}
#### обновить улицу: PUT http://localhost:8080/streets/{id}
пример запроса в body:
{
"name": "Gogolya"
}
#### удалить улицу: DELETE http://localhost:8080/streets/{id}

### Crossroads:
#### получить перекресток: GET http://localhost:8080/crossroads/{id}
#### получить список перекрестков: GET http://localhost:8080/crossroads
#### создать новый перекресток: POST http://localhost:8080/crossroads
пример запроса в body:
{
"street_id": 1,
"city_id": 1
}
#### обновить перекресток: PUT http://localhost:8080/crossroads/{id}
пример запроса в body:
{
"street_id": 1,
"city_id": 1
}
#### удалить перекресток: DELETE http://localhost:8080/crossroads/{id}


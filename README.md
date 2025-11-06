# location-service
The service provides information about **cities**, **streets**, and **crossroads**.

## Entity schema
### City
```
- id
- name
```
### Street
```
- id
- name
- city_id
```
### Crossroad
```
- id
- street_id
```

## Requirements:
The location-service must implement simple **CRUD** operations (```controller```, ```service``` and ```repository``` layers) for all entities.

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

### 2. Клонирование репозитория

```bash
git clone https://github.com/xp10it/location-service.git  
cd location-service
```
### 3. Создать файл с переменными окружения
#### запустить команду в корне проекта:
cp .env.example .env

## 4. Поднять PostgreSQL (через Docker)
###  cборка и запуск контейнеров:

```bash
docker compose build
docker compose up -d
```
### проверить, что контейнеры поднялись: 
docker ps

## 5. Применить миграции
### зайти в контейнер "арр",
docker exec -it location-service sh

### выполнить там команду:
migrate -path internal/db/migrations -database "$DATABASE_URL" up

## 6. Запустить приложение 
### выполнить команду:
go run cmd/location-service/main.go

### после запуска приложение будет доступно по адресу:
http://localhost:8080


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


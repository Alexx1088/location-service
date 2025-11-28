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
<details>
  <summary><del>1) Create the <code>Crossing</code> entity</del></summary>

  ```javascript
  - id
  - crossroad_id
  - event_time
  ```

</details>
<details>
  <summary>
    <del>2) Implement the <code>create</code>  and <code>getAll</code> REST methods.</del>
  </summary>
  The <code>getAll</code> method should include pagination, with <code>from</code> and <code>to</code> 
  query params for filtering by date and a <code>crossroad</code> query param for filtering by 
  <code>crossroad_id</code>
</details>
<details>
  <summary>
    3) Implement sending the <code>crossing event</code> to <code>kafka topic</code> using concurrency
  </summary>
  The main idea is implement saving to the <code>database</code> and sending to <code>kafka topic</code> at the same time.
  This approach needs to consider the possibility of an inconsistent state 
  (e.g., saving to the database succeeds, but sending to kafka fails — in which case, all actions should be rolled back.
  If the database failed, kafka shouldn't receive the event
</details>

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


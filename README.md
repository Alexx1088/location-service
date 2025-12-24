# Location Service

A production-ready backend service for managing urban traffic data, designed with scalability, reliability, and high-load scenarios in mind.

The service is implemented in **Go** and follows clean architectural principles, with a clear separation of concerns between transport, business logic, and persistence layers. It uses **PostgreSQL** for data storage, **Docker** for environment consistency, and **golang-migrate** for controlled schema versioning.

The system exposes **RESTful APIs** and is architected to support **gRPC-based inter-service communication**, making it suitable for microservice environments. The service is designed to be **containerized and deployable in Kubernetes**, enabling horizontal scaling and resilient operation under load.

The platform is also prepared for **asynchronous event processing via Kafka**, allowing reliable real-time data pipelines and event-driven workflows.

## Tech Stack

- Go
- PostgreSQL
- Docker & Docker Compose
- golang-migrate
- REST API
- Kafka 
- gRPC
- Kubernetes

##  Project Setup

### 1.Install Go

Download and install Go from the official website according to your operating system:  
https://go.dev/dl

#### Verify the installation:

go version

#### Example output:
go version go1.24.1 linux/amd64

### 2. Environment Variables
After cloning the repository, create an environment configuration file:
```bash
cp .env.example .env
```
Edit the .env file if needed

## 3. Quick Start (Recommended)
To start the full development environment (PostgreSQL, Kafka, Kafka UI, Minikube, and Go backend), run:
```bash
./dev-start.sh
```
Make sure the script is executable:
```bash
chmod +x dev-start.sh
```
After startup, the application will be available at: http://localhost:8080

## Manual Setup (Optional)
If you prefer to run services manually or debug individual components, follow the steps below.
### Start PostgreSQL (Docker)
####  Build and start containers:

```bash
docker compose build
docker compose up -d
```
### Verify that containers are running: 
docker ps

## 4. Run the Application

```bash
go run cmd/location-service/main.go
```
### After startup, the application will be available at:
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
    <del> 3) Implement sending the <code>crossing event</code> to <code>kafka topic</code> using concurrency
 </del> </summary>
  The main idea is implement saving to the <code>database</code> and sending to <code>kafka topic</code> at the same time.
  This approach needs to consider the possibility of an inconsistent state 
  (e.g., saving to the database succeeds, but sending to kafka fails — in which case, all actions should be rolled back.
  If the database failed, kafka shouldn't receive the event
</details>

## API Endpoints
### Cities:
#### Get city by ID: GET http://localhost:8080/cities/{id}
#### Get list of cities: GET http://localhost:8080/cities
#### Create a city: POST http://localhost:8080/cities
{
"name": "Taraz"
}
#### Update a city: PUT http://localhost:8080/cities/{id}

{
"name": "Taraz"
}
#### Delete a city: DELETE http://localhost:8080/cities/{id}

### Streets:
#### Get street by ID: GET http://localhost:8080/streets/{id}
#### Get list of streets: GET http://localhost:8080/streets
#### Create a street: POST http://localhost:8080/streets

{
"name": "Gogolya"
}
#### Update a street: PUT http://localhost:8080/streets/{id}

{
"name": "Gogolya"
}
#### Delete a street: DELETE http://localhost:8080/streets/{id}

### Crossroads:
#### get crossroad by ID: GET http://localhost:8080/crossroads/{id}
#### Get list of crossroads: GET http://localhost:8080/crossroads
#### Create a crossroad: POST http://localhost:8080/crossroads

{
"street_id": 1,
"city_id": 1
}
#### Update a crossroad: PUT http://localhost:8080/crossroads/{id}

{
"street_id": 1,
"city_id": 1
}
#### Delete a crossroad: DELETE http://localhost:8080/crossroads/{id}


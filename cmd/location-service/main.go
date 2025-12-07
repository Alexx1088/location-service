package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "location-service/cmd/location-service/docs"
	"location-service/internal/db"
	"location-service/internal/kafka"
	"location-service/internal/repository"
	"location-service/internal/router"
	"location-service/internal/worker"
	"log"
	"net/http"
)

// @title Location Service API
// @version 1.0
// @description API for managing cities, streets, and crossroads
// @host localhost:8080
// @BasePath /

func main() {

	db.RunMigrations()
	pool := db.Connect()

	defer pool.Close()

	outboxRepo := repository.NewOutboxRepository(pool)

	producer, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	outboxWorker := worker.NewOutboxWorker(
		outboxRepo,
		producer,
		"crossings.events",
	)

	outboxWorker.Start()

	r := router.NewRouter(pool)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	fmt.Println("Swagger docs: http://localhost:8080/swagger/index.html")

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

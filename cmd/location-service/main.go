package main

import (
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "location-service/cmd/location-service/docs"
	"location-service/internal/config"
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

	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Using Kafka brokers: %v\n", cfg.Kafka.Brokers)

	producer, err := kafka.NewKafkaProducer(cfg.Kafka.Brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workers := cfg.Outbox.WorkersCount

	for i := 0; i < workers; i++ {
		w := worker.NewOutboxWorker(
			outboxRepo,
			producer,
			"crossings.events",
			cfg.Outbox.BatchSize,
			cfg.Outbox.Interval,
		)

		go w.Start(ctx)
	}

	r := router.NewRouter(pool)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	swaggerURL := fmt.Sprintf("http://%s:%s/swagger/index.html", cfg.Server.Host, cfg.Server.Port)
	fmt.Println("Swagger docs:", swaggerURL)

	address := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	fmt.Println("Server running on:", address)
	if err := http.ListenAndServe(address, r); err != nil {
		panic(err)
	}
}

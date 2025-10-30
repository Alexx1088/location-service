package main

import (
	"fmt"
	"location-service/internal/db"
	"location-service/internal/router"
	"net/http"
)

func main() {

	pool := db.Connect()
	defer pool.Close()

	r := router.NewRouter(pool)

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"location-service/internal/db"
	"net/http"
)

func main() {

	pool := db.Connect()
	defer pool.Close()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

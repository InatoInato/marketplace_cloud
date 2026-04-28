package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"marketplace_cloud/internal/db"
	"marketplace_cloud/internal/handler"
	"marketplace_cloud/internal/repository"
	"marketplace_cloud/internal/service"
	"net/http"
)

func main() {
	// Retry the connection to the DB
	var conn *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		conn = db.NewDB() // This should be your connection setup function
		err = conn.Ping()
		if err == nil {
			break
		}
		log.Printf("Unable to connect to DB, retrying... attempt %d", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to DB after several attempts")
	}

	queries := db.New(conn)
	defer conn.Close()

	repo := repository.NewProductRepository(queries)
	service := service.NewProductService(repo)
	handler := handler.NewHandler(service)

	http.HandleFunc("/products", handler.GetAll)
	http.HandleFunc("/products/create", handler.CreateProduct)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Starting server
	fmt.Println("Started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
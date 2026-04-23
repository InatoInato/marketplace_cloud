package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func GetAll() []Product{
	return []Product{
		{ID: 1, Name: "MacBook Air 2020", Description: "M1 16/256GB", Price: 599.99},
		{ID: 2, Name: "iPhone 14 Pro", Description: "6/256GB", Price: 499.99},
		{ID: 3, Name: "Sony PlayStation 5", Description: "1TB HDD", Price: 699.99},
	}
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := GetAll()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	})

	http.ListenAndServe(":8080", nil)
}
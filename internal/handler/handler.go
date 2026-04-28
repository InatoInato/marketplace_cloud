package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"marketplace_cloud/internal/db"
	"marketplace_cloud/internal/service"
	"net/http"
)

type Handler struct {
	service service.ProductService
}

func NewHandler(service service.ProductService) *Handler {
	return &Handler{service: service}
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll(r.Context())
	if err != nil {
		log.Println("GetAll error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// basic validation
	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(
		r.Context(),
		db.CreateProductParams{
			Name: req.Name,
			Description: sql.NullString{
				String: req.Description,
				Valid:  req.Description != "",
			},
			Price: floatToString(req.Price),
		},
	)

	if err != nil {
		log.Println("CreateProduct error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, product)
}

// helpers
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func floatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
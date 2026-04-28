package service

import (
	"context"
	"errors"
	"marketplace_cloud/internal/db"
	"marketplace_cloud/internal/repository"
	"strconv"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]db.Product, error)
	CreateProduct(ctx context.Context, req db.CreateProductParams) (db.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll(ctx context.Context) ([]db.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *productService) CreateProduct(
	ctx context.Context,
	req db.CreateProductParams,
) (db.Product, error) {

	price, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return db.Product{}, errors.New("invalid price format")
	}

	if price < 0 {
		return db.Product{}, errors.New("price cannot be negative")
	}

	return s.repo.Create(ctx, req)
}
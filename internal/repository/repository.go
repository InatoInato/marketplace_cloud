package repository

import (
	"context"
	"marketplace_cloud/internal/db"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]db.Product, error)
	Create(ctx context.Context, p db.CreateProductParams) (db.Product, error)
}

type productRepository struct{
	queries *db.Queries
}

func NewProductRepository(q *db.Queries) *productRepository {
	return &productRepository{queries: q}
}

func (r *productRepository) GetAll(
	ctx context.Context,
) ([]db.Product,error) {

	return r.queries.GetProducts(ctx)
}

func (r *productRepository) Create(
	ctx context.Context,
	p db.CreateProductParams,
)(db.Product,error){

	return r.queries.CreateProduct(ctx,p)
}
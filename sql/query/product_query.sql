-- name: GetProducts :many
SELECT * FROM products;

-- name: CreateProduct :one
INSERT INTO products (name, description, price) 
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
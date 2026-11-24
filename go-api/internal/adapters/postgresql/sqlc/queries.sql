-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProduct :one
SELECT * FROM products where id = $1;
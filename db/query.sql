-- name: CreateUser :one
INSERT INTO users (nome, email, cpf, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET nome = $1, email = $2, cpf = $3
WHERE id = $4
RETURNING *;

-- name: CreateSetor :one
INSERT INTO setores (nome, local)
VALUES ($1, $2)
RETURNING *;

-- name: GetSetorByID :one
SELECT * FROM setores
WHERE id = $1;

-- name: ListSetores :many
SELECT * FROM setores;

-- name: UpdateSetor :one
UPDATE setores
SET nome = $1, local = $2
WHERE id = $3
RETURNING *;

-- name: DeleteSetor :exec
DELETE FROM setores
WHERE id = $1;

-- name: CreateBem :one
INSERT INTO bens (id, nome, status, tipo, setor_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetBemByID :one
SELECT * FROM bens
WHERE id = $1;

-- name: ListBens :many
SELECT * FROM bens;

-- name: UpdateBem :one
UPDATE bens
SET nome = $1, status = $2, tipo = $3, setor_id = $4
WHERE id = $5
RETURNING *;

-- name: DeleteBem :exec
DELETE FROM bens
WHERE id = $1;

-- FORNECEDORES

-- name: CreateFornecedor :one
INSERT INTO fornecedores (nome, cnpj, contato)
VALUES ($1, $2, $3)
RETURNING id, nome, cnpj, contato, created_at;

-- name: ListFornecedores :many
SELECT id, nome, cnpj, contato, created_at
FROM fornecedores;

-- name: GetFornecedorByID :one
SELECT id, nome, cnpj, contato, created_at
FROM fornecedores
WHERE id = $1;

-- name: UpdateFornecedor :one
UPDATE fornecedores
SET nome = $1,
    cnpj = $2,
    contato = $3
WHERE id = $4
RETURNING id, nome, cnpj, contato, created_at;

-- name: DeleteFornecedor :exec
DELETE FROM fornecedores
WHERE id = $1;

-- FABRICANTES

-- name: CreateFabricante :one
INSERT INTO fabricantes (nome, cnpj)
VALUES ($1, $2)
RETURNING *;

-- name: GetFabricanteByID :one
SELECT * FROM fabricantes
WHERE id = $1;

-- name: DeleteFabricante :exec
DELETE FROM fabricantes
WHERE id = $1;

-- name: UpdateFabricante :one
UPDATE fabricantes
SET nome = $1, cnpj = $2
WHERE id = $3
RETURNING *;


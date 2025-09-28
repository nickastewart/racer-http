-- name: GetUserById :one
SELECT id, first_name, last_name, email, created_at FROM user WHERE id = ?;

-- name: CreateUser :one
INSERT INTO user (first_name, last_name, email, password) values (?, ?, ?, ?)
    RETURNING first_name, last_name, email, created_at;

-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email FROM user WHERE email = ?;

-- name: GetUserByEmailForLogin :one
SELECT id, email, password FROM user WHERE email = ?;

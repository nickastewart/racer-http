-- name: GetUser :one
SELECT id, first_name, last_name, email, created_at FROM user WHERE id = ?;

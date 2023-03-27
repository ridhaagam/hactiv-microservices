-- name: GetBooks :many
SELECT * FROM books
ORDER BY id DESC;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: AddBook :one
INSERT INTO books (title, author, descr)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateBook :exec
UPDATE books
  set title = $2,
  author = $3,
  descr = $4
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;
-- +goose Up
CREATE TABLE books (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    authors TEXT NOT NULL,
    pages INT NOT NULL DEFAULT 0,
    isbn TEXT UNIQUE,
    cover_url TEXT,
    published_year INT,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now'))
);

-- +goose StatementBegin
CREATE TRIGGER update_books_updated_at
AFTER UPDATE ON books
FOR EACH ROW
BEGIN
    UPDATE books SET updated_at = (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')) WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS update_books_updated_at;
DROP TABLE IF EXISTS books;

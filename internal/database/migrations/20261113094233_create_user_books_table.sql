-- +goose Up
CREATE TABLE user_books (
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    book_id INTEGER NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    collection TEXT CHECK(collection in ('want_to_read', 'currently_reading', 'backlog', 'completed')) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    PRIMARY KEY (user_id, book_id)
);

-- +goose StatementBegin
CREATE TRIGGER update_user_books_updated_at
AFTER UPDATE ON user_books
FOR EACH ROW
BEGIN
    UPDATE user_books SET updated_at = (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')) WHERE user_id = OLD.user_id AND book_id = OLD.book_id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS update_user_books_updated_at;
DROP TABLE IF EXISTS user_books;

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books
(
    id          serial,
    name        varchar(255) NOT NULL,
    author      varchar(255),
    description varchar(255)
);

CREATE UNIQUE INDEX IF NOT EXISTS books_idx on books(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;

DROP INDEX IF EXISTS books_idx;
-- +goose StatementEnd

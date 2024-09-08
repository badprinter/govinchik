-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE users
ADD is_deleted BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE users
DROP column is_deleted
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE table account
(
    id SERIAL REFERENCES users(id) ON DELETE SET NULL,
    username VARCHAR(32),
    bio_info VARCHAR(4096),
    city varchar(20),
    photo bytea

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS account;
-- +goose StatementEnd

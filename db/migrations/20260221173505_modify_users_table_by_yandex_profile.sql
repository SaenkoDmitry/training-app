-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN yandex_id TEXT UNIQUE,
    ADD COLUMN yandex_login TEXT,
    ADD COLUMN email TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN yandex_id,
    DROP COLUMN yandex_login,
    DROP COLUMN email;
-- +goose StatementEnd

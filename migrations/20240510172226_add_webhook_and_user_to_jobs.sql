-- +goose Up
-- +goose StatementBegin
ALTER TABLE jobs
ADD COLUMN user_id integer,
ADD COLUMN webhook_url VARCHAR(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE jobs
DROP COLUMN user_id,
DROP COLUMN webhook_url;
-- +goose StatementEnd

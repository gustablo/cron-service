-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS jobs (
  uuid           VARCHAR(36) NOT NULL primary key,
  name           VARCHAR(255) NOT NULL,
	execution_time timestamp,
  last_run       timestamp,
	expression     VARCHAR(36) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
-- +goose StatementEnd

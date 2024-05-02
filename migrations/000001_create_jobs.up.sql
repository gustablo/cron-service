BEGIN;

CREATE TABLE IF NOT EXISTS jobs (
  uuid VARCHAR(36) NOT NULL primary key,
	execution_time timestamp,
	expression    VARCHAR(12) NOT NULL,
	name          VARCHAR(255) NOT NULL
);

COMMIT;

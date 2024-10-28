-- migrate:up
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	password BYTEA NOT NULL
);

-- migrate:down
DROP TABLE users;


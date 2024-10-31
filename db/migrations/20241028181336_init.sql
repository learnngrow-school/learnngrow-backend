-- migrate:up
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	email TEXT NOT NULL,
	password BYTEA NOT NULL
);

-- migrate:down
DROP TABLE users;


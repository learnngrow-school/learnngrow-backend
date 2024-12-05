-- migrate:up
CREATE TABLE school_reviews (
	id SERIAL PRIMARY KEY,
	details TEXT NOT NULL,
	author_name TEXT NOT NULL
);

-- migrate:down
DROP TABLE school_reviews;


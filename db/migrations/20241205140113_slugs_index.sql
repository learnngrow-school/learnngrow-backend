-- migrate:up
CREATE UNIQUE INDEX users_slug_idx ON users(slug);

CREATE UNIQUE INDEX courses_slug_idx ON courses(slug);

-- migrate:down
DROP INDEX courses_slug_idx;

DROP INDEX users_slug_idx;


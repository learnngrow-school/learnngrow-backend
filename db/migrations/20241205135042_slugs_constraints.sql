-- migrate:up
ALTER TABLE users
	ALTER COLUMN slug SET NOT NULL;

ALTER TABLE courses
	ALTER COLUMN slug SET NOT NULL;

-- migrate:down
ALTER TABLE courses
	ALTER COLUMN slug DROP NOT NULL;

ALTER TABLE users
	ALTER COLUMN slug DROP NOT NULL;


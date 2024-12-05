-- migrate:up
ALTER TABLE users
	 ADD COLUMN slug TEXT UNIQUE;

ALTER TABLE courses
	 ADD COLUMN slug TEXT UNIQUE;

-- hopefully only for tests
UPDATE users SET slug = substr(md5(random()::text), 27);
UPDATE courses SET slug = substr(md5(random()::text), 27);

-- migrate:down
ALTER TABLE courses
	 DROP COLUMN slug;

ALTER TABLE users
	 DROP COLUMN slug;


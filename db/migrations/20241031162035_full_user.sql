-- migrate:up
ALTER TABLE users
ADD CONSTRAINT email_unique UNIQUE (email),
ADD COLUMN is_teacher BOOLEAN DEFAULT false,
ADD COLUMN is_superuser BOOLEAN DEFAULT false,
ADD COLUMN first_name TEXT NOT NULL DEFAULT 'default_first',
ADD COLUMN middle_name TEXT,
ADD COLUMN last_name TEXT NOT NULL DEFAULT 'default_last';

-- migrate:down
ALTER TABLE users
DROP CONSTRAINT email_unique,
DROP COLUMN is_teacher,
DROP COLUMN is_superuser,
DROP COLUMN first_name,
DROP COLUMN middle_name,
DROP COLUMN last_name;


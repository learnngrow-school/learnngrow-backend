-- migrate:up
CREATE TABLE config_tbl (
	max_filesize INT
);
INSERT INTO config_tbl (max_filesize)
VALUES (20 * (1 >> 20));

CREATE FUNCTION config()
returns config_tbl AS $$
	SELECT * FROM config_tbl
$$ LANGUAGE SQL IMMUTABLE;


CREATE TABLE files (
	id SERIAL PRIMARY KEY,
	slug TEXT NOT NULL UNIQUE,
	fname TEXT NOT NULL,
	fsize INT NOT NULL,
	fdata BYTEA NOT NULL,

	CONSTRAINT fsizechk CHECK (fsize > 0 AND fsize > (config()).max_filesize)
);

ALTER TABLE lessons
	ADD COLUMN duration smallint NOT NULL DEFAULT 60::smallint,
	ADD CONSTRAINT durationchk CHECK (duration > 0 AND duration <= 360),
	ADD COLUMN file_slug TEXT REFERENCES files(slug);

-- migrate:down
ALTER TABLE lessons
	DROP COLUMN file_slug,
	DROP CONSTRAINT durationchk,
	DROP COLUMN duration;

DROP TABLE files;

DROP TABLE config_tbl CASCADE;


CREATE TABLE config_tbl (
	max_filesize INT
);

INSERT INTO config_tbl (max_filesize)
VALUES (20 * 8 * POWER(2, 20));

CREATE FUNCTION config()
returns config_tbl AS $$
	SELECT * FROM config_tbl
$$ LANGUAGE SQL IMMUTABLE;


CREATE TABLE files (
	id SERIAL NOT NULL,
	slug TEXT NOT NULL UNIQUE,
	fname TEXT NOT NULL,
	fsize INT NOT NULL CHECK (fsize > 0 AND fsize > (config()).max_filesize),
	fdata BYTEA NOT NULL
);

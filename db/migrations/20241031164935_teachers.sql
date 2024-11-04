-- migrate:up
CREATE TABLE subjects (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL
);

CREATE TABLE teachers (
	user_id INTEGER PRIMARY KEY,
	subject_ids INTEGER[],
	biography TEXT DEFAULT '',

	CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);

-- migrate:down
DROP TABLE teachers;

DROP TABLE subjects;


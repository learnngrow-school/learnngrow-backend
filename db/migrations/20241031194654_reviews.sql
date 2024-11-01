-- migrate:up
CREATE TABLE teacher_reviews (
	id SERIAL PRIMARY KEY,
	rating SMALLINT NOT NULL,
	details TEXT,
	author_name TEXT NOT NULL,
	teacher_id INTEGER REFERENCES teachers(user_id)
);

CREATE TABLE course_reviews (
	id SERIAL PRIMARY KEY,
	rating SMALLINT NOT NULL,
	details TEXT,
	author_name TEXT NOT NULL,
	course_id INTEGER REFERENCES courses(id)
);

-- migrate:down
DROP TABLE course_reviews;

DROP TABLE teacher_reviews;


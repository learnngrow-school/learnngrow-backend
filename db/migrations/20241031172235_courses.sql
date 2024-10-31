-- migrate:up
CREATE TABLE categories (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL
);

CREATE TABLE courses (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	price INTEGER NOT NULL DEFAULT 0,
	year SMALLINT,
	category_id INTEGER NOT NULL REFERENCES categories(id),
	subject_id INTEGER NOT NULL REFERENCES subjects(id)
);

CREATE TABLE courses_teachers (
	course_id INTEGER NOT NULL REFERENCES courses(id),
	teacher_id INTEGER NOT NULL REFERENCES teachers(user_id),
	PRIMARY KEY (course_id, teacher_id)
);

-- migrate:down
DROP TABLE courses_teachers;

DROP TABLE courses;

DROP TABLE categories;


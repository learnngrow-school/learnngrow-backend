-- migrate:up
CREATE TABLE lessons (
	id SERIAL PRIMARY KEY,
	student_id INT NOT NULL REFERENCES users(id),
	teacher_id INT NOT NULL REFERENCES teachers(user_id),
	timestamp_m INT NOT NULL,
	teacher_notes TEXT NOT NULL,
	homework TEXT NOT NULL
);

-- migrate:down
DROP TABLE lessons;


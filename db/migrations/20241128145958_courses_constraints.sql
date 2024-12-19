-- migrate:up
ALTER TABLE teacher_reviews
	ALTER COLUMN details SET NOT NULL,
	ALTER COLUMN teacher_id SET NOT NULL;

ALTER TABLE course_reviews
	ALTER COLUMN details SET NOT NULL,
	ALTER COLUMN course_id SET NOT NULL;

-- migrate:down
ALTER TABLE teacher_reviews
	ALTER COLUMN details DROP NOT NULL,
	ALTER COLUMN teacher_id DROP NOT NULL;

ALTER TABLE course_reviews
	ALTER COLUMN details DROP NOT NULL,
	ALTER COLUMN course_id DROP NOT NULL;


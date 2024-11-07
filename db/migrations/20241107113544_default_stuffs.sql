-- migrate:up
INSERT INTO categories (title)
VALUES ('default_category');

INSERT INTO subjects (title)
VALUES ('default_subject');

-- migrate:down
DELETE FROM categories WHERE title = 'default_category';
DELETE FROM subjects WHERE title = 'default_subject';


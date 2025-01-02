-- migrate:up
ALTER TABLE lessons
	ADD COLUMN classwork TEXT NOT NULL DEFAULT '';

ALTER TABLE users
	ADD COLUMN tg_chat_id TEXT NOT NULL DEFAULT '';

-- migrate:down
ALTER TABLE lessons
	DROP COLUMN classwork;

ALTER TABLE users
	DROP COLUMN tg_chat_id;


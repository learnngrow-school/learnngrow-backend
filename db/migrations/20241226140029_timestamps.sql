-- migrate:up
ALTER TABLE lessons
	ADD COLUMN ts TIMESTAMPTZ;

UPDATE lessons
	SET ts = to_timestamp(timestamp_m * 60);

ALTER TABLE lessons
	DROP COLUMN timestamp_m,
	ALTER COLUMN ts SET NOT NULL;

-- migrate:down
ALTER TABLE lessons
	ADD COLUMN timestamp_m INT AFTER teacher_id;

UPDATE lessons
	SET timestamp_m = EXTRACT(EPOCH FROM ts)::integer / 60;

ALTER TABLE lessons
	DROP COLUMN ts,
	ALTER COLUMN timestamp_m SET NOT NULL;


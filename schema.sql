DROP TABLE IF EXISTS  users, tasks, lables, tasks_lables;

CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE tasks(
	id SERIAL PRIMARY KEY,
	opened BIGINT NOT NULL DEFAULT extract(epoch from now()),
	closed BIGINT DEFAULT 0 ,
	author_id BIGINT REFERENCES users(id) DEFAULT 0,
	assigned_id BIGINT REFERENCES users(id) DEFAULT 0,
	title TEXT,
	content TEXT
);

CREATE TABLE lables(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE tasks_lables(
	task_id BIGINT REFERENCES tasks(id),
	label_id BIGINT REFERENCES lables(id)
);

INSERT INTO users(id, name) VALUES(0, 'default');

SELECT * FROM users;
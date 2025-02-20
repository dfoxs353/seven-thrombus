CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    roles JSONB NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS tokens(
    user_id INT4 NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    expires_in TIMESTAMP NOT NULL,

    PRIMARY KEY(user_id, refresh_token)
);

CREATE TABLE IF NOT EXISTS study_groups(
	id SERIAL PRIMARY KEY,
	course int4 NOT NULL,
	title varchar(255) NOT NULL,
	created_at int4 NOT NULL
);

CREATE TABLE IF NOT EXISTS student_groups(
	id SERIAL PRIMARY KEY,
	student_id INT4 NOT NULL REFERENCES users(id),
	group_id int4 NOT NULL REFERENCES study_groups(id),
	created_at int4 NOT NULL 
);

CREATE TABLE IF NOT EXISTS disciplines(
	id SERIAL PRIMARY KEY,
	title varchar(255) NOT NULL,
	created_at int4 NOT NULL 
);

CREATE TABLE IF NOT EXISTS files(
	id SERIAL PRIMARY KEY,
	"path" varchar(255) NOT NULL,
	src_name varchar(255) NOT NULL 
);

CREATE TABLE IF NOT EXISTS lessons(
	id SERIAL PRIMARY KEY,
	discipline_id int4 NOT NULL REFERENCES disciplines(id),
	teacher_id INT4 NOT NULL REFERENCES users(id),
	group_id int4 NOT NULL REFERENCES study_groups(id),
	"date" TIMESTAMPTZ NOT NULL, 
	is_online boolean NOT NULL DEFAULT false,
	file_ids JSON NOT NULL DEFAULT '[]',
	created_at int4 NOT NULL 
);

CREATE TABLE IF NOT EXISTS attendance_logs(
	lesson_id int4 REFERENCES users(id),
	student_id INT4 NOT NULL REFERENCES users(id),
	is_present boolean NOT NULL,
	created_at int4 NOT NULL,
	
	PRIMARY KEY (lesson_id, student_id)
);

CREATE TABLE IF NOT EXISTS tasks(
	id SERIAL PRIMARY KEY,
	title varchar(255) NOT NULL,
	description varchar(255) NOT NULL,
	group_id int4 NOT NULL REFERENCES study_groups(id),
	deadline date NOT NULL,
	file_ids JSON NOT NULL DEFAULT '[]',
	discipline_id int4 NOT NULL REFERENCES disciplines(id),
	teacher_id INT4 NOT NULL REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS solutions(
	id SERIAL PRIMARY KEY,
	student_id INT4 NOT NULL REFERENCES users(id),
	task_id INT4 NOT NULL REFERENCES tasks(id),
	file_ids JSON NOT NULL DEFAULT '[]',
	mark int4 NOT NULL
);
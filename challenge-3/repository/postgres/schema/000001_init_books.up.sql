CREATE TABLE books (
	id BIGSERIAL PRIMARY KEY,
	Title text NOT NULL,
	Author text NOT NULL,
	Descr text NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);

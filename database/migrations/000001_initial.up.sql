BEGIN;

-- users
CREATE TABLE IF NOT EXISTS users (
  id SERIAL NOT NULL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


--  posts
CREATE TABLE IF NOT EXISTS posts (
	id SERIAL NOT NULL PRIMARY KEY,
	user_id INTEGER NOT NULL REFERENCES users(id),
	contents TEXT NOT NULL,
	like_count   INTEGER NOT NULL DEFAULT 0,
	deleted boolean NOT NULL DEFAULT FALSE,
  
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX posts_created_at_idx ON posts(created_at DESC, id) WHERE deleted = FALSE;


--  post attachments
CREATE TABLE IF NOT EXISTS post_attachments (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	post_id INTEGER NOT NULL REFERENCES posts(id),
	content_type TEXT NOT NULL,
	attachment_size INT NOT NULL,
	attachment_data BYTEA NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX post_attachments_post_id_created_at_idx ON post_attachments(post_id, created_at);


COMMIT;

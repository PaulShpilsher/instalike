BEGIN;

-- users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);


--  posts
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER NOT NULL REFERENCES users(id),
    body TEXT NOT NULL,
    like_count   INTEGER NOT NULL DEFAULT 0,
    deleted boolean NOT NULL DEFAULT FALSE
);

CREATE INDEX posts_created_at_idx ON posts(created_at DESC, id) WHERE deleted = FALSE;


--  post attachments
CREATE TABLE IF NOT EXISTS post_attachments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    post_id INTEGER NOT NULL REFERENCES posts(id),
    content_type TEXT NOT NULL,
    attachment_size INT NOT NULL,
    attachment_data BYTEA NOT NULL
);

CREATE INDEX post_attachments_post_id_created_at_idx ON post_attachments(post_id, created_at);

-- posts_view 
CREATE OR REPLACE VIEW posts_view
AS 
SELECT      p.id, 
            p.created_at,
            p.updated_at,
            u.email AS author, 
            p.body,
            p.like_count, 
            ARRAY(SELECT id FROM post_attachments WHERE post_id = p.id) AS attachment_ids,
            (p.created_at <> p.updated_at) AS updated
FROM 		posts AS p
INNER JOIN	users AS u
        ON	p.user_id = u.id
WHERE       deleted IS FALSE
;


COMMIT;

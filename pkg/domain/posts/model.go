package posts

import (
	"time"
)

type Post struct {
	Id      int       `db:"id" json:"id"`
	Created time.Time `db:"created_at" json:"created"`
	Updated time.Time `db:"updated_at" json:"updated"`

	Author        string `db:"author" json:"author"`
	Body          string `db:"body" json:"body"`
	LikeCount     int    `db:"like_count" json:"likeCount"`
	IsUpdated     bool   `db:"updated" json:"IsUpdated"`
	AttachmentIds string `db:"attachment_ids" json:"attachmentIds"`
}

type PostAttachment struct {
	Id          int    `db:"id" json:"id"`
	ContentType string `db:"content_type" json:"contentType"`
	Size        int    `db:"attachment_size" json:"size"`
}

type PostComment struct {
	Id        int       `db:"id" json:"id"`
	Created   time.Time `db:"created_at" json:"created"`
	Updated   time.Time `db:"updated_at" json:"updated"`
	Author    string    `db:"author" json:"author"`
	Body      string    `db:"body" json:"body"`
	IsUpdated bool      `db:"updated" json:"IsUpdated"`
}

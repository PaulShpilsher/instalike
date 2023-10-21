package posts

import "time"

type Post struct {
	Id      int       `db:"id" json:"id"`
	Created time.Time `db:"created_at" json:"created"`
	Updated time.Time `db:"updated_at" json:"updated"`

	UserId    int    `db:"user_id" json:"userId"`
	Contents  string `db:"contents" json:"contents"`
	LikeCount int    `db:"like_count" json:"likeCount"`
}

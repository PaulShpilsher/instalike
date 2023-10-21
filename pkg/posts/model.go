package posts

import "time"

type Post struct {
	Id      int
	Created time.Time `db:"created_at" json:"created_at"`
	Updated time.Time `db:"updated_at" json:"updated_at"`
}

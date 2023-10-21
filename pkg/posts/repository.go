package posts

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	*sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreatePost(userId int, contents string) (Post, error) {

	post := Post{}

	if err := r.DB.Get(&post, "INSERT INTO posts (user_id, contents) VALUES($1, $2) RETURNING *", userId, contents); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return Post{}, err
	}

	return post, nil
}

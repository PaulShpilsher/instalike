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

func (r *repository) CreatePost(userId int, contents string) (int, error) {
	var id int
	if err := r.DB.Get(&id, "INSERT INTO posts (user_id, contents) VALUES($1, $2) RETURNING id", userId, contents); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return 0, err
	}

	return id, nil
}

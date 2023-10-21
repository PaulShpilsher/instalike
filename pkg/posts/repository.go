package posts

import (
	"fmt"
	"log"
	"strings"

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

	if err := r.DB.Get(&post, "INSERT INTO posts (user_id, contents) VALUES($1, $2) RETURNING id, user_id, contents, like_count, created_at, updated_at", userId, contents); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return Post{}, err
	}

	return post, nil
}

func (r *repository) GetPosts() ([]Post, error) {

	posts := []Post{}

	if err := r.DB.Select(&posts, "SELECT id, user_id, contents, like_count, created_at, updated_at FROM posts ORDER BY created_at DESC"); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return []Post{}, err
	}

	return posts, nil
}

func (r *repository) GetPostById(postId int) (Post, error) {

	post := Post{}

	if err := r.DB.Get(&post, "SELECT id, user_id, contents, like_count, created_at, updated_at FROM posts WHERE id = $1 LIMIT 1", postId); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return Post{}, fmt.Errorf("post not found")
		}
		log.Printf("[DB ERROR]: %v", err)
		return Post{}, err
	}

	return post, nil
}

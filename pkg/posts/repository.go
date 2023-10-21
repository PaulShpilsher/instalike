package posts

import (
	"log"
	"strings"

	"github.com/PaulShpilsher/instalike/pkg/utils"
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

	sql := `
		INSERT INTO posts (user_id, contents)
		VALUES($1, $2)
		RETURNING id, user_id, contents, like_count, created_at, updated_at
	`

	post := Post{}
	if err := r.DB.Get(&post, sql, userId, contents); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return Post{}, err
	}

	return post, nil
}

func (r *repository) GetPosts() ([]Post, error) {

	sql := `
		SELECT id, user_id, contents, like_count, created_at, updated_at
		FROM posts
		WHERE deleted IS FALSE
		ORDER BY created_at DESC
		`

	posts := []Post{}
	if err := r.DB.Select(&posts, sql); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return []Post{}, err
	}

	return posts, nil
}

func (r *repository) GetPost(postId int) (Post, error) {

	sql := `
		SELECT	id,user_id, contents, like_count, created_at, updated_at
		FROM	posts
		WHERE	id = $1
			AND	deleted IS FALSE
		LIMIT 1
	`
	post := Post{}
	if err := r.DB.Get(&post, sql, postId); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return Post{}, utils.ErrNotFound
		}
		log.Printf("[DB ERROR]: %v", err)
		return Post{}, err
	}

	return post, nil
}

func (r *repository) GetAuthor(postId int) (int, error) {

	sql := `
		SELECT	user_id
		FROM	posts
		WHERE	id = $1 AND	deleted IS FALSE
		LIMIT 1
	`
	var authorId int
	if err := r.DB.Get(&authorId, sql, postId); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return 0, utils.ErrNotFound
		}
		log.Printf("[DB ERROR]: %v", err)
		return 0, err
	}

	return authorId, nil
}

func (r *repository) DeletePost(postId int) error {

	// we dont delete actual data from the database
	// instead we just set the deleted flag to true
	sql := `
		UPDATE	posts
		SET		deleted = TRUE
		WHERE	id = $1 AND	deleted IS FALSE
	`
	if result, err := r.DB.Exec(sql, postId); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return err
	} else if rows, _ := result.RowsAffected(); rows == 0 {
		return utils.ErrNotFound
	}

	return nil
}

func (r *repository) UpdatePost(postId int, contents string) error {
	sql := `
		UPDATE	posts 
		SET		contents = $2
		WHERE	id = $1 AND deleted IS FALSE
	`
	if result, err := r.DB.Exec(sql, postId, contents); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return err
	} else if rows, _ := result.RowsAffected(); rows == 0 {
		return utils.ErrNotFound
	}

	return nil
}

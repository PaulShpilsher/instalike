package posts

import (
	"log"
	"strings"

	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type postsRepository struct {
	*sqlx.DB
}

func NewPostsRepository(db *sqlx.DB) *postsRepository {
	return &postsRepository{
		DB: db,
	}
}

func (r *postsRepository) CreatePost(userId int, body string) (int, error) {

	sql := `
		INSERT INTO posts (user_id, body)
		VALUES($1, $2)
		RETURNING id
	`

	var postId int
	if err := r.DB.Get(&postId, sql, userId, body); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return 0, err
	}

	return postId, nil
}

func (r *postsRepository) GetPosts() ([]Post, error) {

	sql := `
		SELECT	id, created_at, updated_at, author, body, like_count, updated, attachment_ids
		FROM posts_view
		ORDER BY created_at DESC
		`

	posts := []Post{}
	if err := r.DB.Select(&posts, sql); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return []Post{}, err
	}

	return posts, nil
}

func (r *postsRepository) GetPost(postId int) (Post, error) {

	sql := `
		SELECT	id, created_at, updated_at, author, body, like_count, updated, attachment_ids
		FROM	posts_view
		WHERE	id = $1
		LIMIT	1
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

func (r *postsRepository) GetAuthor(postId int) (int, error) {

	sql := `
		SELECT	user_id
		FROM	posts
		WHERE	id = $1
		AND		deleted IS FALSE
		LIMIT	1
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

func (r *postsRepository) DeletePost(postId int) error {

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

func (r *postsRepository) UpdatePost(postId int, body string) error {
	sql := `
		UPDATE	posts 
		SET		body = $2,
				updated_at = CURRENT_TIMESTAMP
		WHERE	id = $1 AND deleted IS FALSE
	`
	if result, err := r.DB.Exec(sql, postId, body); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return err
	} else if rows, _ := result.RowsAffected(); rows == 0 {
		return utils.ErrNotFound
	}

	return nil
}

package posts

import (
	"log"
	"strings"

	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/jmoiron/sqlx"
)

///
// PostCommentsRepository - post multimedia Comments data store
//

type postCommentsRepository struct {
	*sqlx.DB
}

func NewPostCommentsRepository(db *sqlx.DB) *postCommentsRepository {
	return &postCommentsRepository{
		DB: db,
	}
}

func (r *postCommentsRepository) CreateComment(postId int, userId int, body string) error {

	sql := `
		INSERT INTO post_comments (post_id, author_id, body) VALUES($1, $2, $3)
	`
	if _, err := r.DB.Exec(sql, postId, userId, body); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return utils.ErrNotFound
		}
		return err
	}

	return nil
}

func (r *postCommentsRepository) GetComments(postId int) ([]PostComment, error) {
	sql := `
		SELECT id, created_at, updated_at, author, body, updated
		FROM post_comments_view
		WHERE post_id = $1
		ORDER BY created_at DESC
		`

	var comments []PostComment
	if err := r.DB.Select(&comments, sql, postId); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return []PostComment{}, err
	}

	return comments, nil
}

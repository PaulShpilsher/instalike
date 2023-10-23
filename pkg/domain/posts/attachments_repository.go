package posts

import (
	"log"

	"github.com/jmoiron/sqlx"
)

//
// PostAttachmentsRepository - post multimedia attachments data store
//

type postAttachmentsRepository struct {
	*sqlx.DB
}

func NewPostAttachmentsRepository(db *sqlx.DB) *postAttachmentsRepository {
	return &postAttachmentsRepository{
		DB: db,
	}
}

func (r *postAttachmentsRepository) CreatePostAttachment(postId int, contentType string, binary []byte) error {

	sql := `
		INSERT INTO post_attachments (post_id, content_type, attachment_size, attachment_data) 
		VALUES($1, $2, $3, $4)
		RETURNING id;
	`
	var attachmentId int64
	if err := r.DB.Get(&attachmentId, sql, postId, contentType, len(binary), binary); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return err
	}

	return nil
}

func (r *postAttachmentsRepository) GetPostAttachments(postId int) ([]PostAttachment, error) {

	sql := `
		SELECT	id, content_type, attachment_size
		FROM	post_attachments
		WHERE 	post_id = $1
	`

	attachments := []PostAttachment{}
	if err := r.DB.Select(&attachments, sql); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return []PostAttachment{}, err
	}

	return attachments, nil
}

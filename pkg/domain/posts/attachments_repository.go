package posts

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type attachmentsRepository struct {
	*sqlx.DB
}

func NewAttachmentRepository(db *sqlx.DB) *attachmentsRepository {
	return &attachmentsRepository{
		DB: db,
	}
}

func (r *attachmentsRepository) CreatePostAttachment(postId int, contentType string, binary []byte) error {

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

func (r *attachmentsRepository) GetPostAttachments(postId int) ([]Attachment, error) {

	sql := `
		SELECT	id, content_type, attachment_size
		FROM	post_attachments
		WHERE 	post_id = $1
	`

	attachments := []Attachment{}
	if err := r.DB.Select(&attachments, sql); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return []Attachment{}, err
	}

	return attachments, nil
}

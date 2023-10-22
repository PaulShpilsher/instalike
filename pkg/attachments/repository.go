package attachments

import (
	"log"
	"strings"

	"github.com/PaulShpilsher/instalike/pkg/utils"
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

func (r *attachmentsRepository) GetAttachment(attachmentId int) (Attachment, error) {

	sql := `
		SELECT	content_type, attachment_size, attachment_data
		FROM	post_attachments
		WHERE 	id = $1
		LIMIT 	1
	`

	attachment := Attachment{}
	if err := r.DB.Get(&attachment, sql, attachmentId); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return Attachment{}, utils.ErrNotFound
		}
		log.Printf("[DB ERROR]: %v", err)
		return Attachment{}, err
	}

	return attachment, nil
}

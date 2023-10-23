package media

import (
	"log"
	"strings"

	"github.com/PaulShpilsher/instalike/pkg/utils"
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

func (r *postAttachmentsRepository) GetPostAttachment(attachmentId int) (MultimediaData, error) {

	sql := `
		SELECT	content_type, attachment_size, attachment_data
		FROM	post_attachments
		WHERE 	id = $1
		LIMIT 	1
	`

	attachment := MultimediaData{}
	if err := r.DB.Get(&attachment, sql, attachmentId); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return MultimediaData{}, utils.ErrNotFound
		}
		log.Printf("[DB ERROR]: %v", err)
		return MultimediaData{}, err
	}

	return attachment, nil
}

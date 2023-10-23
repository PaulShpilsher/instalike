package media

// MediaService interface - multimedia business logic
type MediaService interface {
	GetPostAttachment(attachmentId int) (MultimediaData, error)
}

// PostAttachmentsRepository interface - post multimedia attachments data store
type PostAttachmentsRepository interface {
	GetPostAttachment(attachmentId int) (MultimediaData, error)
}

package media

type MediaService interface {
	GetAttachment(attachmentId int) (Attachment, error)
}

type AttachmentRepository interface {
	GetAttachment(attachmentId int) (Attachment, error)
}

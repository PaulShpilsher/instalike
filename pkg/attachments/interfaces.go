package attachments

// AttachmentsService interface declares users business logic
type AttachmentsService interface {
	GetAttachment(attachmentId int) (Attachment, error)
}

type AttachmentRepository interface {
	GetAttachment(attachmentId int) (Attachment, error)
}

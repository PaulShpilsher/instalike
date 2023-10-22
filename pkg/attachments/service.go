package attachments

type attachmentService struct {
	attachmentRepo AttachmentRepository
}

func NewAttachmentService(attachmentRepo AttachmentRepository) *attachmentService {
	return &attachmentService{
		attachmentRepo: attachmentRepo,
	}
}

func (s *attachmentService) GetAttachment(attachmentId int) (Attachment, error) {
	return s.attachmentRepo.GetAttachment(attachmentId)
}

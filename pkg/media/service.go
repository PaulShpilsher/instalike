package media

type mediaService struct {
	attachmentRepo AttachmentRepository
}

func NewMediaService(attachmentRepo AttachmentRepository) *mediaService {
	return &mediaService{
		attachmentRepo: attachmentRepo,
	}
}

func (s *mediaService) GetAttachment(attachmentId int) (Attachment, error) {
	return s.attachmentRepo.GetAttachment(attachmentId)
}

package media

// MediaService - multimedia business logic
type mediaService struct {
	postAttachmentRepo PostAttachmentsRepository
}

func NewMediaService(attachmentRepo PostAttachmentsRepository) *mediaService {
	return &mediaService{
		postAttachmentRepo: attachmentRepo,
	}
}

func (s *mediaService) GetPostAttachment(attachmentId int) (MultimediaData, error) {
	return s.postAttachmentRepo.GetPostAttachment(attachmentId)
}

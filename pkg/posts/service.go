package posts

import (
	"io"

	"github.com/PaulShpilsher/instalike/pkg/utils"
)

type service struct {
	postsRepo      PostsRepository
	attachmentRepo AttachmentRepository
}

func NewService(postsRepo PostsRepository, attachmentRepo AttachmentRepository) *service {
	return &service{
		postsRepo:      postsRepo,
		attachmentRepo: attachmentRepo,
	}
}

func (s *service) CreatePost(userId int, contents string) (Post, error) {
	post, err := s.postsRepo.CreatePost(userId, contents)
	return post, err
}

func (s *service) GetPosts() ([]Post, error) {
	posts, err := s.postsRepo.GetPosts()
	return posts, err
}

func (s *service) GetPost(postId int) (Post, error) {
	post, err := s.postsRepo.GetPost(postId)
	return post, err
}

func (s *service) DeletePost(userId int, postId int) error {

	if err := s.validatePostAuthor(userId, postId); err != nil {
		return err
	}

	return s.postsRepo.DeletePost(postId)
}

func (s *service) UpdatePost(userId int, postId int, contents string) error {

	if err := s.validatePostAuthor(userId, postId); err != nil {
		return err
	}

	return s.postsRepo.UpdatePost(postId, contents)
}

func (s *service) AttachFileToPost(userId int, postId int, contentType string, size int, reader io.Reader) error {

	err := s.validatePostAuthor(userId, postId)
	if err != nil {
		return err
	}

	// TODO: read data in chunks, but for now just read the whole file in memory
	// and store it to the database
	binary, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = s.attachmentRepo.CreatePostAttachment(postId, contentType, binary)
	return err
}

// private functions

func (s *service) validatePostAuthor(userId int, postId int) error {
	authorId, err := s.postsRepo.GetAuthor(postId)
	if err != nil {
		return err
	}

	if userId != authorId {
		return utils.ErrForbidden
	}

	return nil
}

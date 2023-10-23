package posts

import (
	"io"

	"github.com/PaulShpilsher/instalike/pkg/utils"
)

//
// PostsService - posts business logic
//

type postsService struct {
	postsRepo      PostsRepository
	attachmentRepo PostAttachmentsRepository
}

func NewPostsService(postsRepo PostsRepository, attachmentRepo PostAttachmentsRepository) *postsService {
	return &postsService{
		postsRepo:      postsRepo,
		attachmentRepo: attachmentRepo,
	}
}

func (s *postsService) CreatePost(userId int, body string) (int, error) {
	postId, err := s.postsRepo.CreatePost(userId, body)
	return postId, err
}

func (s *postsService) GetPosts() ([]Post, error) {
	posts, err := s.postsRepo.GetPosts()
	return posts, err
}

func (s *postsService) GetPost(postId int) (Post, error) {
	post, err := s.postsRepo.GetPost(postId)
	return post, err
}

func (s *postsService) DeletePost(userId int, postId int) error {

	if err := s.validatePostAuthor(userId, postId); err != nil {
		return err
	}

	return s.postsRepo.DeletePost(postId)
}

func (s *postsService) UpdatePost(userId int, postId int, body string) error {

	if err := s.validatePostAuthor(userId, postId); err != nil {
		return err
	}

	return s.postsRepo.UpdatePost(postId, body)
}

func (s *postsService) AttachFileToPost(userId int, postId int, contentType string, size int, reader io.Reader) error {

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

func (s *postsService) LikePost(userId int, postId int) error {

	if liked, err := s.postsRepo.DidUserLikePost(postId, userId); err != nil {
		return err
	} else if liked {
		return utils.ErrAlreadyExists // user already liked this post
	}

	return s.postsRepo.LikePost(postId, userId)
}

// private functions

func (s *postsService) validatePostAuthor(userId int, postId int) error {
	authorId, err := s.postsRepo.GetAuthor(postId)
	if err != nil {
		return err
	}

	if userId != authorId {
		return utils.ErrForbidden
	}

	return nil
}

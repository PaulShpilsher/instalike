package posts

import "github.com/PaulShpilsher/instalike/pkg/utils"

type service struct {
	repo PostsRepository
}

func NewService(repo PostsRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreatePost(userId int, contents string) (Post, error) {
	post, err := s.repo.CreatePost(userId, contents)
	return post, err
}

func (s *service) GetPosts() ([]Post, error) {
	posts, err := s.repo.GetPosts()
	return posts, err
}

func (s *service) GetPost(postId int) (Post, error) {
	post, err := s.repo.GetPost(postId)
	return post, err
}

func (s *service) DeletePost(userId int, postId int) error {

	if err := s.validatePostAuthor(userId, postId); err != nil {
		return err
	}

	return s.repo.DeletePost(postId)
}

func (s *service) UpdatePost(userId int, postId int, contents string) error {

	if err := s.validatePostAuthor(userId, postId); err != nil {
		return err
	}

	return s.repo.UpdatePost(postId, contents)
}

// private functions

func (s *service) validatePostAuthor(userId int, postId int) error {
	authorId, err := s.repo.GetAuthor(postId)
	if err != nil {
		return err
	}

	if userId != authorId {
		return utils.ErrForbidden
	}

	return nil
}

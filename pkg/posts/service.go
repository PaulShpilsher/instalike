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

func (s *service) GetPostById(postId int) (Post, error) {
	post, err := s.repo.GetPostById(postId)
	return post, err
}

func (s *service) DeletePostById(userId int, postId int) error {

	authorId, err := s.repo.GetAuthor(postId)
	if err != nil {
		return err
	}

	if userId != authorId {
		return utils.ErrUnauthorized
	}

	return s.repo.DeletePostById(postId)
}

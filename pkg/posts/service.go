package posts

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

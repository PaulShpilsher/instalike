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

func (s *service) GetPosts() ([]Post, error) {
	posts, err := s.repo.GetPosts()
	return posts, err
}

func (s *service) GetPostById(postId int) (Post, error) {
	post, err := s.repo.GetPostById(postId)
	return post, err
}

func (s *service) DeletePostById(postId int) error {
	return s.repo.DeletePostById(postId)
}

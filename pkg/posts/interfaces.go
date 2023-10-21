package posts

// PostsService interface declares users business logic
type PostsService interface {
	CreatePost(userId int, contents string) (Post, error)
	GetPosts() ([]Post, error)
	GetPostById(postId int) (Post, error)
	UpdatePost(userId int, postId int, contents string) error
	DeletePostById(userId int, postId int) error
}

// PostsRepository interface declares users data store logic
type PostsRepository interface {
	GetAuthor(postId int) (int, error)
	CreatePost(userId int, contents string) (Post, error)
	UpdatePost(postId int, contents string) error
	GetPosts() ([]Post, error)
	GetPostById(postId int) (Post, error)
	DeletePostById(postId int) error
}

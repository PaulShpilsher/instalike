package posts

// PostsService interface declares users business logic
type PostsService interface {
	CreatePost(userId int, contents string) (Post, error)
}

// PostsRepository interface declares users data store logic
type PostsRepository interface {
	CreatePost(userId int, contents string) (Post, error)
}

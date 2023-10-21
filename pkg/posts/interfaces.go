package posts

import "io"

// PostsService interface declares users business logic
type PostsService interface {
	CreatePost(userId int, contents string) (Post, error)
	GetPosts() ([]Post, error)
	GetPost(postId int) (Post, error)
	UpdatePost(userId int, postId int, contents string) error
	DeletePost(userId int, postId int) error
	AttachFileToPost(userId int, postId int, contentType string, size int, reader io.Reader) error
}

// PostsRepository interface declares users data store logic
type PostsRepository interface {
	GetAuthor(postId int) (int, error)
	CreatePost(userId int, contents string) (Post, error)
	UpdatePost(postId int, contents string) error
	DeletePost(postId int) error
	GetPosts() ([]Post, error)
	GetPost(postId int) (Post, error)
	AttachFileToPost(postId int, contentType string, binary []byte) error
}

package posts

import "io"

// PostsService interface - posts business logic
type PostsService interface {
	CreatePost(userId int, body string) (int, error)
	GetPosts() ([]Post, error)
	GetPost(postId int) (Post, error)
	UpdatePost(userId int, postId int, body string) error
	DeletePost(userId int, postId int) error
	AttachFileToPost(userId int, postId int, contentType string, size int, reader io.Reader) error
}

// PostsRepository interface - posts data store logic
type PostsRepository interface {
	GetAuthor(postId int) (int, error)
	CreatePost(userId int, body string) (int, error)
	UpdatePost(postId int, body string) error
	DeletePost(postId int) error
	GetPosts() ([]Post, error)
	GetPost(postId int) (Post, error)
}

// PostAttachmentsRepository interface - post multimedia attachments data store
type PostAttachmentsRepository interface {
	CreatePostAttachment(postId int, contentType string, binary []byte) error
}

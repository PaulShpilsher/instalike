package posts

import (
	"time"

	"github.com/PaulShpilsher/instalike/pkg/utils"
)

type CreatePostInput struct {
	Body string `json:"body" validate:"required"`
}

type CreatePostOutput struct {
	Id int `json:"id"`
}

type UpdatePostInput struct {
	Body string `json:"body" validate:"required"`
}

type GetPostOutput struct {
	Id            int       `json:"id"`
	Created       time.Time `json:"created"`
	IsUpdated     bool      `json:"isUpdated"`
	Author        string    `json:"author"`
	Body          string    `json:"body"`
	LikeCount     int       `json:"likeCount"`
	AttachmentIds []int64   `json:"attachmentIds"`
}

// factory functions

func MakeGetPostOutput(post Post) GetPostOutput {
	return GetPostOutput{
		Id:            post.Id,
		Created:       post.Created,
		IsUpdated:     post.IsUpdated,
		Author:        post.Author,
		Body:          post.Body,
		LikeCount:     post.LikeCount,
		AttachmentIds: utils.ParseStringToInt64Array(post.AttachmentIds),
	}
}

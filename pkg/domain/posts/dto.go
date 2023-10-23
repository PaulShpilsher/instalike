package posts

import (
	"time"

	"github.com/PaulShpilsher/instalike/pkg/utils"
)

type createPostInput struct {
	Body string `json:"body" validate:"required"`
}

type createPostOutput struct {
	Id int `json:"id"`
}

type updatePostInput struct {
	Body string `json:"body" validate:"required"`
}

type getPostOutput struct {
	Id            int       `json:"id"`
	Created       time.Time `json:"created"`
	IsUpdated     bool      `json:"isUpdated"`
	Author        string    `json:"author"`
	Body          string    `json:"body"`
	LikeCount     int       `json:"likeCount"`
	AttachmentIds []int64   `json:"attachmentIds"`
}

// factory functions

func makeGetPostOutput(post Post) getPostOutput {
	return getPostOutput{
		Id:            post.Id,
		Created:       post.Created,
		IsUpdated:     post.IsUpdated,
		Author:        post.Author,
		Body:          post.Body,
		LikeCount:     post.LikeCount,
		AttachmentIds: utils.ParseStringToInt64Array(post.AttachmentIds),
	}
}

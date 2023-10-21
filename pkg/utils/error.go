package utils

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type ErrorOutput struct {
	Message string           `json:"message,omitempty"`
	Errors  []*ErrorResponse `json:"errors,omitempty"`
}

func NewValidationErrorOutput(errors []*ErrorResponse) *ErrorOutput {
	return &ErrorOutput{
		Errors: errors,
	}
}

func NewErrorOutput(message string) *ErrorOutput {
	return &ErrorOutput{
		Message: message,
	}
}

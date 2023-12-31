package utils

import "errors"

//
// Error related helper
//

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrForbidden     = errors.New("forbidden")
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

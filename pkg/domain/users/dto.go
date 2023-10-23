package users

import "time"

type registerInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=5"`
}

type loginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginOutput struct {
	Token string `json:"token"`
}

type userOutput struct {
	UserId  int       `json:"userId"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

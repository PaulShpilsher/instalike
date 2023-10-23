package users

import "time"

type User struct {
	Id           int
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Created      time.Time `db:"created_at" json:"createdAt"`
	Updated      time.Time `db:"updated_at" json:"updatedAt"`
}

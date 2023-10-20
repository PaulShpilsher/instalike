package users

import (
	"log"
	"math"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	*sqlx.DB
}

func NewRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(email string, passwordHash string) (int, error) {
	var id int
	if err := r.DB.Get(&id, "INSERT INTO users (email, password_hash) VALUES($1, $2) RETURNING id", email, passwordHash); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return math.MinInt, err
	}

	return id, nil
}

func (r *userRepository) GetUserById(id int) (User, error) {
	var user User
	if err := r.DB.Get(&user, "SELECT id, email, created_at, updated_at FROM users WHERE id = $1 LIMIT 1", id); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (User, error) {
	var user User
	if err := r.DB.Get(&user, "SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = $1 LIMIT 1", email); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return User{}, err
	}

	return user, nil
}

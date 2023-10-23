package users

import (
	"log"
	"strings"

	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/jmoiron/sqlx"
)

//
// UsersRepository - users data store logic
//

type usersRepository struct {
	*sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *usersRepository {
	return &usersRepository{
		DB: db,
	}
}

func (r *usersRepository) CreateUser(email string, passwordHash string) (int, error) {
	var id int
	if err := r.DB.Get(&id, "INSERT INTO users (email, password_hash) VALUES($1, $2) RETURNING id", email, passwordHash); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return 0, utils.ErrAlreadyExists
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (r *usersRepository) GetUserById(id int) (User, error) {
	var user User
	if err := r.DB.Get(&user, "SELECT id, email, created_at, updated_at FROM users WHERE id = $1 LIMIT 1", id); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return User{}, err
	}

	return user, nil
}

func (r *usersRepository) GetUserByEmail(email string) (User, error) {
	var user User
	if err := r.DB.Get(&user, "SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = $1 LIMIT 1", email); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return User{}, err
	}

	return user, nil
}

func (r *usersRepository) GetUserExistsByEmail(email string) (bool, error) {
	var exists bool
	if err := r.DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1) AS exists", email); err != nil {
		log.Printf("[DB ERROR]: %v", err)
		return false, err
	}

	return exists, nil
}

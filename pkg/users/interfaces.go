package users

// UserService interface declares users business logic
type UserService interface {
	Register(email string, password string) (userId int, err error)
	Login(email string, password string) (userId int, token string, err error)
	GetUserById(id int) (user User, err error)
}

// UserRepository interface declares users data store logic
type UserRepository interface {
	CreateUser(email string, passwordHash string) (int, error)
	GetUserById(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserExistsByEmail(email string) (bool, error)
}

package users

// UsersService interface - users business logic
type UsersService interface {
	Register(email string, password string) (userId int, err error)
	Login(email string, password string) (userId int, token string, err error)
	GetUserById(id int) (user User, err error)
}

// UsersRepository interface - users data store logic
type UsersRepository interface {
	CreateUser(email string, passwordHash string) (int, error)
	GetUserById(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserExistsByEmail(email string) (bool, error)
}

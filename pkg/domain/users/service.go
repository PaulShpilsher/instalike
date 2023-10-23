package users

import (
	"strconv"

	"github.com/PaulShpilsher/instalike/pkg/token"
	"github.com/PaulShpilsher/instalike/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

// UsersService - users business logic
type usersService struct {
	usersRepo  UsersRepository
	jwtService token.JwtService
}

func NewService(usersRepo UsersRepository, jwtService token.JwtService) *usersService {
	return &usersService{
		usersRepo:  usersRepo,
		jwtService: jwtService,
	}
}

func (s *usersService) Register(email string, password string) (int, error) {

	userExists, err := s.usersRepo.GetUserExistsByEmail(email)
	if err != nil {
		return 0, err
	}
	if userExists {
		return 0, utils.ErrAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userId, err := s.usersRepo.CreateUser(email, string(passwordHash))
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *usersService) Login(email string, password string) (int, string, error) {

	user, err := s.usersRepo.GetUserByEmail(email)
	if err != nil {
		return 0, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return 0, "", err
	}

	token, err := s.jwtService.CreateToken(strconv.Itoa((user.Id)))
	if err != nil {
		return 0, "", err
	}

	return user.Id, token, nil
}

func (s *usersService) GetUserById(id int) (User, error) {

	user, err := s.usersRepo.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

package users

import (
	"fmt"
	"strconv"

	"github.com/PaulShpilsher/instalike/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo UserRepository
	jwt  token.JwtService
}

func NewService(repo UserRepository, jwt token.JwtService) *service {
	return &service{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *service) Register(email string, password string) (int, error) {

	userExists, err := s.repo.GetUserExistsByEmail(email)
	if err != nil {
		return 0, err
	}
	if userExists {
		return 0, fmt.Errorf("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userId, err := s.repo.CreateUser(email, string(passwordHash))
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *service) Login(email string, password string) (int, string, error) {

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return 0, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return 0, "", err
	}

	token, err := s.jwt.CreateToken(strconv.Itoa((user.Id)))
	if err != nil {
		return 0, "", err
	}

	return user.Id, token, nil
}

func (s *service) GetUserById(id int) (User, error) {

	user, err := s.repo.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

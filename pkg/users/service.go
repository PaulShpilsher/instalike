package users

import (
	"math"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(email string, passwordHash string) (int, error)
	GetUserById(id int) (User, error)
	GetUserByEmail(email string) (User, error)
}

type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(email string, password string) (int, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return math.MinInt, err
	}

	userId, err := s.repo.CreateUser(email, string(passwordHash))
	if err != nil {
		return math.MinInt, err
	}

	return userId, nil
}

func (s *service) Login(email string, password string) (int, error) {

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return math.MinInt, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return math.MinInt, err
	}

	return user.Id, nil
}

func (s *service) GetUserById(id int) (User, error) {

	user, err := s.repo.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

package users

import (
	"log"
	"math"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	// repo ArticlesRepository
	dummy string
}

func NewService() *service {
	return &service{
		dummy: "dummy",
	}
}

func (s *service) Signup(email string, password string) (int, error) {
	log.Println("Service.Signup", email, password)
	_, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return math.MinInt, err
	}

	return 1, nil
}

func (s *service) Login(email string, password string) (userId int, err error) {
	log.Println("Service.Login", email, password)

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordHash)); err != nil {
		if err != nil {
			return math.MinInt, err
		}
	}

	return 1, nil
}

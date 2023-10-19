package users

import "log"

type service struct {
	// repo ArticlesRepository
	dummy string
}

func NewService() *service {
	return &service{
		dummy: "dummy",
	}
}

func (s *service) Signup(email string, password string) (userId int, token string, err error) {
	log.Println("Service.Signup", email, password)
	return 1, "token", nil
}

func (s *service) Login(email string, password string) (userId int, token string, err error) {
	log.Println("Service.Login", email, password)
	return 1, "token", nil
}

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

func (s *service) Signup(email string, password string) (userId int, err error) {
	log.Println("Service.Signup", email, password)
	return 1, nil
}

func (s *service) Login(email string, password string) (userId int, err error) {
	log.Println("Service.Login", email, password)
	return 1, nil
}

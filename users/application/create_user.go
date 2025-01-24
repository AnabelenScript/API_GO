package application

import (
	"fmt"
	"api_prueba/users/domain"
)

type UserService struct {
	Repo domain.UserRepository
}

func (s *UserService) CreateUser(name, email string) error {
	user := &domain.User{Name: name, Email: email}
	fmt.Println("Si funca")
	return s.Repo.Save(user)
}

package application

import (
	"API_GO/users/domain/entities"
	"API_GO/users/domain"
)

type CreateUser struct {
	Repo domain.UserRepository
}

func NewCreateUser(repo domain.UserRepository) *CreateUser {
	return &CreateUser{Repo: repo}
}

func (uc *CreateUser) Execute(name, email string, user_type int, password string) error {
	user := &entities.User{Name: name, Email: email, User_type: user_type, Password: password}
	return uc.Repo.Save(user)
}

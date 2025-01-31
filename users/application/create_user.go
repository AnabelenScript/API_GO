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

func (uc *CreateUser) Execute(name, email string) error {
	user := &entities.User{Name: name, Email: email}
	return uc.Repo.Save(user)
}

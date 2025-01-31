package application

import (
	"API_GO/users/domain"
	"API_GO/users/domain/entities"
)

type GetAllUsers struct {
	Repo domain.UserRepository
}

func NewGetAllUsers(repo domain.UserRepository) *GetAllUsers {
	return &GetAllUsers{Repo: repo}
}

func (uc *GetAllUsers) Execute() ([]*entities.User, error) {
	return uc.Repo.GetAll()
}
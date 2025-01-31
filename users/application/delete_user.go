package application

import (
	"API_GO/users/domain"
	"errors"
)

type DeleteUser struct {
	Repo domain.UserRepository
	
}

func NewDeleteUser(repo domain.UserRepository) *DeleteUser {
	return &DeleteUser{Repo: repo}
}

func (uc *DeleteUser) Execute(id uint) error {
	user, err := uc.Repo.FindByID(id)
	if err != nil {
		return errors.New("usuario no encontrado")
	}
	return uc.Repo.Delete(uint(user.ID))
}

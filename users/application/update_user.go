package application

import (
	"API_GO/users/domain"
)

type UpdateUser struct {
	Repo domain.UserRepository
}

func NewUpdateUser(repo domain.UserRepository) *UpdateUser {
	return &UpdateUser{Repo: repo}
}

func (uc *UpdateUser) Execute(id uint, name, email string) error {
	user, err := uc.Repo.FindByID(id) 
	if err != nil {
		return err 
	}
	user.Name = name
	user.Email = email
	return uc.Repo.Update(user)
}

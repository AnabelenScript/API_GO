package application

import(
	"API_GO/users/domain/entities"
	"API_GO/users/domain"
)

type GetLastUser struct {
	Repo domain.UserRepository
}

func NewGetLastUser(repo domain.UserRepository) *GetLastUser {
	return &GetLastUser{Repo: repo}
}

func (uc *GetLastUser) GetLastUserProcess() (*entities.User, error) {
	return uc.Repo.GetLastAddedUser()
}

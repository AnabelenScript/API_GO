package application

import (
	"API_GO/users/domain"
)

type LoginUser struct {
	Repo domain.UserRepository
}

func NewLoginUser(repo domain.UserRepository) *LoginUser {
	return &LoginUser{Repo: repo}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
}
func (uc *LoginUser) Execute(email, password string) (*LoginResponse, error) {
	user, err := uc.Repo.Login(email, password)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		UserType: user.User_type,
	}, nil
}

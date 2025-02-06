package domain

import (
	"API_GO/users/domain/entities"
)

type UserRepository interface {
	Save(user *entities.User) error
	FindByID(id uint) (*entities.User, error) 
	Update(user *entities.User) error 
	Delete(id uint) error
	GetAll() ([]*entities.User, error)
	GetLastAddedUser()(*entities.User, error)
}

package domain

import "API_GO/users/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
}

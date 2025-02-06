package domain

import "API_GO/desserts/domain/entities"

type DessertRepository interface {
	Save(dessert *entities.Dessert) error
	FindByID(id uint) (*entities.Dessert, error) 
	Update(dessert *entities.Dessert) error
	Delete(id uint) error
	GetAll() ([]*entities.Dessert, error)
	GetLastDessert()(*entities.Dessert, error)
	GetAllDessertsForDelete()([]* entities.Dessert, error)
}

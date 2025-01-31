package application

import (
	"API_GO/desserts/domain/entities"
	"API_GO/desserts/domain"
)

type CreateDessert struct {
	Repo domain.DessertRepository
}

func NewCreateDessert(repo domain.DessertRepository) *CreateDessert {
	return &CreateDessert{Repo: repo}
}

func (uc *CreateDessert) Execute(name string, flavor string, price int, quantity int) error {
	dessert := &entities.Dessert{Name: name, Flavor: flavor, Price: price, Quantity: quantity}
	return uc.Repo.Save(dessert)
}

package application

import (
	"API_GO/desserts/domain"
)

type UpdateDessert struct {
	Repo domain.DessertRepository
}

func NewUpdateDessert(repo domain.DessertRepository) *UpdateDessert {
	return &UpdateDessert{Repo: repo}
}

func (uc *UpdateDessert) Execute(id uint, name, flavor string, price int, quantity int) error {
	dessert, err := uc.Repo.FindByID(id) 
	if err != nil {
		return err 
	}
	dessert.Name = name
	dessert.Flavor = flavor
	dessert.Price = price
	dessert.Quantity= quantity
	return uc.Repo.Update(dessert)
}

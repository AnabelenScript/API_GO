package application

import (
	"API_GO/desserts/domain"
	"errors"
)

type DeleteDessert struct {
	Repo domain.DessertRepository
	
}

func NewDeleteDessert(repo domain.DessertRepository) *DeleteDessert {
	return &DeleteDessert{Repo: repo}
}

func (uc *DeleteDessert) Execute(id uint) error {
	dessert, err := uc.Repo.FindByID(id)
	if err != nil {
		return errors.New("Postre no encontrado")
	}
	return uc.Repo.Delete(uint(dessert.Id))
}

package application

import (
	"API_GO/desserts/domain"
	"API_GO/desserts/domain/entities"
	"errors"
)

type GetPriceDessert struct {
	Repo domain.DessertRepository
}

func NewGetPriceDessert(repo domain.DessertRepository) *GetPriceDessert {
	return &GetPriceDessert{Repo: repo}
}

func (uc *GetPriceDessert) Execute(price uint) ([]*entities.Dessert, error) {
	desserts, err := uc.Repo.FindByPrice(price)
	if err != nil {
		return nil, errors.New("Postre no encontrado")
	}

	return desserts, nil
}
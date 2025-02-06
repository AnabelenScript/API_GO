package application

import (
	"API_GO/desserts/domain/entities"
	"API_GO/desserts/domain"
)

type GetDeletedDessert struct {
	Repo domain.DessertRepository
}

func NewGetDeletedDessert(repo domain.DessertRepository) *GetDeletedDessert {
	return &GetDeletedDessert{Repo: repo}
}

func (uc *GetDeletedDessert) GetAllDessertsForDeleteProcess() ([]*entities.Dessert, error) {
	return uc.Repo.GetAllDessertsForDelete()
}

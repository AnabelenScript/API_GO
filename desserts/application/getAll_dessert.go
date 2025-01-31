package application

import (
	"API_GO/desserts/domain"
	"API_GO/desserts/domain/entities"
)

type GetAllDesserts struct {
	Repo domain.DessertRepository
}

func NewGetAllDesserts(repo domain.DessertRepository) *GetAllDesserts {
	return &GetAllDesserts{Repo: repo}
}

func (uc *GetAllDesserts) Execute() ([]*entities.Dessert, error) {
	return uc.Repo.GetAll()
}
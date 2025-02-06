package application

import(
	"API_GO/desserts/domain/entities"
	"API_GO/desserts/domain"
)

type GetLastDesserts struct {
	Repo domain.DessertRepository
}

func NewGetLastDessert(repo domain.DessertRepository) *GetLastDesserts {
	return &GetLastDesserts{Repo: repo}
}

func (uc *GetLastDesserts) GetLastDessertProcess() (*entities.Dessert, error) {
	return uc.Repo.GetLastDessert()
}

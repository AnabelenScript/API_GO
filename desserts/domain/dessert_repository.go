package domain

import "API_GO/desserts/domain/entities"

type DessertRepository interface {
	Save(dessert *entities.Dessert) error
	
}

package db


import (
	"database/sql"
	"API_GO/desserts/domain"
	"API_GO/desserts/domain/entities"
	"log"
)

type MySQLDessertRepository struct {
	DB *sql.DB
}

func NewMySQLDessertRepository(db *sql.DB) domain.DessertRepository {
	return &MySQLDessertRepository{DB: db}
}


func (r *MySQLDessertRepository) Save(dessert *entities.Dessert) error {
	query := "INSERT INTO dessert (name, flavor, price, quantity) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, dessert.Name, dessert.Flavor, dessert.Price, dessert.Quantity)
	if err != nil {
		log.Printf("Error al agrgear el postre: %v", err)
	}
	return err
}

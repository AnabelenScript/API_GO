package db


import (
	"database/sql"
	"API_GO/desserts/domain"
	"API_GO/desserts/domain/entities"
	"log"
	"errors"
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

func (r *MySQLDessertRepository) FindByID(id uint) (*entities.Dessert, error) {
	query := "SELECT ID, name, flavor, price, quantity FROM dessert WHERE ID = ?"
	row := r.DB.QueryRow(query, id)

	var dessert entities.Dessert
	err := row.Scan(&dessert.Id, &dessert.Name, &dessert.Flavor, &dessert.Price, &dessert.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Postre no encontrado :C")
		}
		log.Printf("Error al buscar postre: %v", err)
		return nil, err
	}
	return &dessert, nil
}

func (r *MySQLDessertRepository) Update(dessert *entities.Dessert) error {
	query := "UPDATE dessert SET name = ?, flavor = ?, price = ?, quantity = ? WHERE ID = ?"
	result, err := r.DB.Exec(query, dessert.Name, dessert.Flavor, dessert.Price, dessert.Quantity, dessert.Id)
	if err != nil {
		log.Printf("Error al actualizar el postre: %v", err)
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no se encontró el postre para actualizar :(")
	}

	return nil
}

func (r *MySQLDessertRepository) Delete(dessertID uint) error {
	query := "DELETE FROM dessert WHERE ID = ?"
	result, err := r.DB.Exec(query, dessertID)
	if err != nil {
		log.Printf("Error al eliminar el postre: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no se encontró el postre para eliminar")
	}

	return nil
}

func (r *MySQLDessertRepository) GetAll() ([]*entities.Dessert, error) {
	query := "SELECT ID, name, flavor, price, quantity FROM dessert"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Error al obtener todos los postres: %v", err)
		return nil, err
	}
	defer rows.Close()

	var desserts []*entities.Dessert
	for rows.Next() {
		dessert := &entities.Dessert{}
		if err := rows.Scan(&dessert.Id, &dessert.Name, &dessert.Flavor,&dessert.Price, &dessert.Quantity); err != nil {
			log.Printf("Error al escanear el postre: %v", err)
			return nil, err
		}
		desserts = append(desserts, dessert)
	}

	return desserts, nil
}

func (r *MySQLDessertRepository) GetLastDessert() (*entities.Dessert, error) {
	rows, err := r.DB.Query("SELECT ID, name, flavor, price, quantity FROM dessert ORDER BY ID DESC LIMIT 1")
	if err != nil {
		log.Printf("Error al obtener el último postre registrado: %v", err)
		return nil, err
	}
	defer rows.Close()

	var dessert entities.Dessert

	if rows.Next() {
		if err := rows.Scan(&dessert.Id, &dessert.Name, &dessert.Flavor, &dessert.Price, &dessert.Quantity); err != nil {
			log.Printf("Error al escanear el postre: %v", err)
			return nil, err
		}
		log.Printf("Último postre agregado -> ID: %d, Nombre: %s, Sabor: %s, Precio: %d, Cantidad: %d", dessert.Id, dessert.Name, dessert.Flavor, dessert.Price, dessert.Quantity)
	} else {
		return nil, errors.New("No se encontraron postres en la base de datos")
	}

	return &dessert, nil
}

func (r *MySQLDessertRepository) GetAllDessertsForDelete() ([]*entities.Dessert, error) {
	query := "SELECT ID, name, flavor, price, quantity FROM dessert"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Error al obtener todos los postres: %v", err)
		return nil, err
	}
	defer rows.Close()

	var desserts []*entities.Dessert
	for rows.Next() {
		dessert := &entities.Dessert{}
		if err := rows.Scan(&dessert.Id, &dessert.Name, &dessert.Flavor,&dessert.Price, &dessert.Quantity); err != nil {
			log.Printf("Error al escanear el postre: %v", err)
			return nil, err
		}
		desserts = append(desserts, dessert)
	}

	return desserts, nil
}


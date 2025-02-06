package db

import (
	"database/sql"
	"API_GO/users/domain"
	"API_GO/users/domain/entities"
	"errors"
	"log"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) domain.UserRepository {
	return &MySQLUserRepository{DB: db}
}


func (r *MySQLUserRepository) Save(user *entities.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := r.DB.Exec(query, user.Name, user.Email)
	if err != nil {
		log.Printf("Error al insertar usuario: %v", err)
	}
	return err
}

func (r *MySQLUserRepository) FindByID(id uint) (*entities.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("usuario no encontrado")
		}
		log.Printf("Error al buscar usuario: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) Update(user *entities.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	result, err := r.DB.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		log.Printf("Error al actualizar usuario: %v", err)
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no se encontró el usuario para actualizar")
	}
	return nil
}

func (r *MySQLUserRepository) Delete(userID uint) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := r.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Error al eliminar usuario: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no se encontró el usuario para eliminar")
	}

	return nil
}

func (r *MySQLUserRepository) GetAll() ([]*entities.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Error al obtener usuarios: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Printf("Error al escanear usuario: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *MySQLUserRepository) GetLastAddedUser() (*entities.User, error) {
	rows, err := r.DB.Query("SELECT ID, name, email FROM users ORDER BY ID DESC LIMIT 1")
	if err != nil {
		log.Printf("Error al obtener el último usuario registrado: %v", err)
		return nil, err
	}
	defer rows.Close()

	var user entities.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Printf("Error al escanear el postre: %v", err)
			return nil, err
		}
		log.Printf("Último usuario agregado -> ID: %d, Nombre: %s, Email: %s", user.ID, user.Name, user.Email)
	} else {
		return nil, errors.New("No se encontraron usuarios en la base de datos")
	}

	return &user, nil
}



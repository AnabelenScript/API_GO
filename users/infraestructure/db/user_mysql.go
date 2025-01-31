package db

import (
	"database/sql"
	"API_GO/users/domain/entities"
	"API_GO/users/domain"
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

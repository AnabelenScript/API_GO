package db

import (
	"database/sql"
	"api_prueba/users/domain"
)
//Repositorio de mysql

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{DB: db}
}

func (r *MySQLUserRepository) Save(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := r.DB.Exec(query, user.Name, user.Email)
	return err
}

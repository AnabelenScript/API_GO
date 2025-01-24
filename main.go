package main

import (
	"api_prueba/helpers"
	"api_prueba/users/application"
	"api_prueba/users/infraestructure/db"
	"api_prueba/users/infraestructure/http_handlers"
	"log"
	"net/http"
)

func main() {
	// Conexión a MySQL
	dbConn := helpers.ConnectToMySQL()
	defer dbConn.Close()

	//Configura el repositorio, servicio y http
	userRepo := db.NewMySQLUserRepository(dbConn)
	userService := &application.UserService{Repo: userRepo}
	userHandler := &http_handlers.UserHandler{Service: userService}

	http.HandleFunc("/users", userHandler.CreateUser)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
    "API_GO/helpers"
    "API_GO/users/application"
    "API_GO/users/infraestructure/db"
    "API_GO/users/infraestructure/controllers"
    "API_GO/users/infraestructure/routes"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
   //Inicia conexion a mysql
    dbConn := helpers.ConnectToMySQL()
    defer dbConn.Close()
    //Configurar repositorio, servicio y controlador
    userRepo := db.NewMySQLUserRepository(dbConn)
	//casos de uso
    createUser := application.NewCreateUser(userRepo)
	updateUser := application.NewUpdateUser(userRepo)
	deleteUser := application.NewDeleteUser(userRepo)
    getAllUser := application.NewGetAllUsers(userRepo)
	//controladores
    createUserController := controllers.NewCreateUserController(createUser)
	updateUserController := controllers.NewUpdateUserController(updateUser)
	deleteUserController := controllers.NewDeleteUserController(deleteUser)
    getAllUserController := controllers.NewGetAllUserController(getAllUser)
    r := gin.Default()
    routes.SetupUsersRoutes(r, createUserController, updateUserController,deleteUserController, getAllUserController)
    log.Println("Server running on :8080")
    r.Run(":8080")
}

package main

import (
    "API_GO/helpers"
    "API_GO/users/application"
    "API_GO/users/infraestructure/db"
    "API_GO/users/infraestructure/controllers"
    "API_GO/users/infraestructure/routes"
    "github.com/gin-gonic/gin"
    "log"

    dessertApplication "API_GO/desserts/application"
    dessertInfra "API_GO/desserts/infraestructure/db"
    dessertControllers "API_GO/desserts/infraestructure/controllers"
    dessertRoutes "API_GO/desserts/infraestructure/routes"
)

func main() {
   //Inicia conexion a mysql
    dbConn := helpers.ConnectToMySQL()
    defer dbConn.Close()
    //Configurar repositorio, servicio y controlador
    userRepo := db.NewMySQLUserRepository(dbConn)
    dessertRepo := dessertInfra.NewMySQLDessertRepository(dbConn)
	//casos de uso
    createUser := application.NewCreateUser(userRepo)
	updateUser := application.NewUpdateUser(userRepo)
	deleteUser := application.NewDeleteUser(userRepo)
    getAllUser := application.NewGetAllUsers(userRepo)
    createDessert := dessertApplication.NewCreateDessert(dessertRepo)
    updatateDessert := dessertApplication.NewUpdateDessert(dessertRepo)
    deleteDessert := dessertApplication.NewDeleteDessert(dessertRepo)
    getAllDessert := dessertApplication.NewGetAllDesserts(dessertRepo)
	//controladores
    createUserController := controllers.NewCreateUserController(createUser)
	updateUserController := controllers.NewUpdateUserController(updateUser)
	deleteUserController := controllers.NewDeleteUserController(deleteUser)
    getAllUserController := controllers.NewGetAllUserController(getAllUser)
    createDessertController := dessertControllers.NewCreateDessertController(createDessert)
    updateDessertController := dessertControllers.NewUpdateUserController(updatateDessert)
    deleteDessertController := dessertControllers.NewDeleteUserController(deleteDessert)
    getAllDessertController := dessertControllers.NewGetAllDessertController(getAllDessert)
    r := gin.Default()
    routes.SetupUsersRoutes(r, createUserController, updateUserController,deleteUserController, getAllUserController)
    dessertRoutes.SetupDessertsRoutes(r, createDessertController, updateDessertController, deleteDessertController, getAllDessertController)
    log.Println("Server running on :8080")
    r.Run(":8080")
}

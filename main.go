package main

import (
	"API_GO/helpers"
	"API_GO/users/application"
	"API_GO/users/infraestructure/controllers"
	"API_GO/users/infraestructure/db"
	"API_GO/users/infraestructure/routes"
	"log"
	/*"time"*/
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	dessertApplication "API_GO/desserts/application"
	dessertControllers "API_GO/desserts/infraestructure/controllers"
	dessertInfra "API_GO/desserts/infraestructure/db"
	dessertRoutes "API_GO/desserts/infraestructure/routes"
)

func main() {
	// Inicia conexión a MySQL
	dbConn := helpers.ConnectToMySQL()
	defer dbConn.Close()

	// Configurar repositorio, servicio y controlador
	userRepo := db.NewMySQLUserRepository(dbConn)
	dessertRepo := dessertInfra.NewMySQLDessertRepository(dbConn)

	// Casos de uso
	createUser := application.NewCreateUser(userRepo)
	updateUser := application.NewUpdateUser(userRepo)
	deleteUser := application.NewDeleteUser(userRepo)
	getAllUser := application.NewGetAllUsers(userRepo)
	/*getLastUserAdded := application.NewGetLastUser(userRepo)*/

	createDessert := dessertApplication.NewCreateDessert(dessertRepo)
	updateDessert := dessertApplication.NewUpdateDessert(dessertRepo)
	deleteDessert := dessertApplication.NewDeleteDessert(dessertRepo)
	getAllDessert := dessertApplication.NewGetAllDesserts(dessertRepo)
	/*getLastDessert := dessertApplication.NewGetLastDessert(dessertRepo)*/
	/*getDeletedDessert := dessertApplication.NewGetDeletedDessert(dessertRepo)*/
	getPriceDessert := dessertApplication.NewGetPriceDessert(dessertRepo)

	// Controladores
	createUserController := controllers.NewCreateUserController(createUser)
	updateUserController := controllers.NewUpdateUserController(updateUser)
	deleteUserController := controllers.NewDeleteUserController(deleteUser)
	getAllUserController := controllers.NewGetAllUserController(getAllUser)
	/*getLastUSerController := controllers.NewGetLastUserController(*getLastUserAdded)*/

	createDessertController := dessertControllers.NewCreateDessertController(createDessert)
	updateDessertController := dessertControllers.NewUpdateUserController(updateDessert)
	deleteDessertController := dessertControllers.NewDeleteUserController(deleteDessert)
	getAllDessertController := dessertControllers.NewGetAllDessertController(getAllDessert)
	/*getLastDesserts := dessertControllers.NewGetLastDessertController(*getLastDessert)*/
	/*getDeletedDessertController := dessertControllers.NewGetDeletedDessertController(*getDeletedDessert)*/
	getPriceDessertController := dessertControllers.NewGetPriceDessertController(getPriceDessert)

	/*go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			getLastDessert.GetLastDessertProcess()
			getLastUserAdded.GetLastUserProcess()
			getDeletedDessert.Repo.GetAllDessertsForDelete()
		}
	}()*/

	r := gin.Default()

	r.Use(cors.Default())

	routes.SetupUsersRoutes(r, createUserController, updateUserController, deleteUserController, getAllUserController, /*getLastUSerController*/)
	dessertRoutes.SetupDessertsRoutes(r, createDessertController, updateDessertController, deleteDessertController, getAllDessertController, /*getLastDesserts, getDeletedDessertController,*/ getPriceDessertController)

	log.Println("Server running on :8080")
	r.Run(":8080")
}

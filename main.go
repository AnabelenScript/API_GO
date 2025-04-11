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
	
	pedidosApplication "API_GO/pedidos/application"
	pedidosControllers "API_GO/pedidos/infraestructure/controllers"
	pedidosInfra "API_GO/pedidos/infraestructure/db"
	pedidosRoutes "API_GO/pedidos/infraestructure/routes"

)

func main() {
	dbConn := helpers.ConnectToMySQL()
	defer dbConn.Close()
	rabbitConn, ch, err := helpers.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()
	defer ch.Close()

	// Configurar repositorio, servicio y controlador
	userRepo := db.NewMySQLUserRepository(dbConn)
	dessertRepo := dessertInfra.NewMySQLDessertRepository(dbConn)
	pedidosRepo := pedidosInfra.NewMySQLPedidosRepository(dbConn, ch)
	// Casos de uso
	createUser := application.NewCreateUser(userRepo)
	updateUser := application.NewUpdateUser(userRepo)
	deleteUser := application.NewDeleteUser(userRepo)
	getAllUser := application.NewGetAllUsers(userRepo)
	loginUser := application.NewLoginUser(userRepo)
	/*getLastUserAdded := application.NewGetLastUser(userRepo)*/

	createDessert := dessertApplication.NewCreateDessert(dessertRepo)
	updateDessert := dessertApplication.NewUpdateDessert(dessertRepo)
	deleteDessert := dessertApplication.NewDeleteDessert(dessertRepo)
	getAllDessert := dessertApplication.NewGetAllDesserts(dessertRepo)
	/*getLastDessert := dessertApplication.NewGetLastDessert(dessertRepo)*/
	/*getDeletedDessert := dessertApplication.NewGetDeletedDessert(dessertRepo)*/
	getPriceDessert := dessertApplication.NewGetPriceDessert(dessertRepo)

	createPedidos := pedidosApplication.NewCreatePedidos(pedidosRepo)
	reducirInventario := pedidosApplication.NewDecreaseStock(pedidosRepo)
	deletePedidos := pedidosApplication.NewDeletePedido(pedidosRepo)
	updatePedidos := pedidosApplication.NewUpdatePedidos(pedidosRepo)
	getAllPedidos := pedidosApplication.NewGetAllPedidos(pedidosRepo)

	// Controladores
	createUserController := controllers.NewCreateUserController(createUser)
	updateUserController := controllers.NewUpdateUserController(updateUser)
	deleteUserController := controllers.NewDeleteUserController(deleteUser)
	getAllUserController := controllers.NewGetAllUserController(getAllUser)
	loginUserController := controllers.NewLoginUserController(loginUser)
	/*getLastUSerController := controllers.NewGetLastUserController(*getLastUserAdded)*/

	createDessertController := dessertControllers.NewCreateDessertController(createDessert)
	updateDessertController := dessertControllers.NewUpdateUserController(updateDessert)
	deleteDessertController := dessertControllers.NewDeleteUserController(deleteDessert)
	getAllDessertController := dessertControllers.NewGetAllDessertController(getAllDessert)
	/*getLastDesserts := dessertControllers.NewGetLastDessertController(*getLastDessert)*/
	/*getDeletedDessertController := dessertControllers.NewGetDeletedDessertController(*getDeletedDessert)*/
	getPriceDessertController := dessertControllers.NewGetPriceDessertController(getPriceDessert)

	createPedidosController := pedidosControllers.NewCreatePedidosController(createPedidos, reducirInventario)
	deletePedidosController := pedidosControllers.NewDeleteUserController(deletePedidos)
	updatePedidosController := pedidosControllers.NewUpdateUserController(updatePedidos)
	getAllPedidosController := pedidosControllers.NewGetAllDessertController(getAllPedidos)
	reducirInventarioController := pedidosControllers.NewDecreaseStockController(reducirInventario)


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

	routes.SetupUsersRoutes(r, createUserController, updateUserController, deleteUserController, getAllUserController, /*getLastUSerController*/ loginUserController)
	dessertRoutes.SetupDessertsRoutes(r, createDessertController, updateDessertController, deleteDessertController, getAllDessertController, /*getLastDesserts, getDeletedDessertController,*/ getPriceDessertController)
	pedidosRoutes.SetupPedidosRoutes(r, createPedidosController, updatePedidosController,  deletePedidosController, getAllPedidosController, reducirInventarioController)
	log.Println("Server running on :8080")
	r.Run(":8080")
}

package routes

import(
	"API_GO/desserts/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDessertsRoutes(
	r *gin.Engine,
	NewCreateDessertController *controllers.CreateDessertController,
	NewUpdateDessertController *controllers.UpdateDessertController,
	NewDeleteDessertController *controllers.DeleteDessertController,
	NewGetAllDessertController *controllers.GetAllDessertController){
	r.POST("/desserts", NewCreateDessertController.Execute)
	r.PUT("/desserts/:id", NewUpdateDessertController.Execute)
	r.DELETE("/desserts/:id", NewDeleteDessertController.Execute)
	r.GET("/desserts", NewGetAllDessertController.Execute)
}
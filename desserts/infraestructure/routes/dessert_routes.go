package routes

import(
	"API_GO/desserts/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDessertsRoutes(
	r *gin.Engine,
	NewCreateDessertController *controllers.CreateDessertController){
	r.POST("/desserts", NewCreateDessertController.Execute)
}
package routes

import(
	"API_GO/pedidos/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPedidosRoutes(
	r *gin.Engine,
	NewCreatePedidosController *controllers.CreatePedidosController,
	NewUpdatePedidosController *controllers.UpdatePedidosController,
	NewDeletePedidosController *controllers.DeletePedidoController,
	NewGetAllPedidosController *controllers.GetAllPedidosController,
	NewDecreaseStockController *controllers.DecreaseStockController ){
	r.POST("/pedidos", NewCreatePedidosController.Execute)
	r.PUT("/pedidos/:id", NewUpdatePedidosController.Execute)
	r.PUT("/pedidos/inventario", NewDecreaseStockController.Execute)
	r.DELETE("/pedidos/:id", NewDeletePedidosController.Execute)
	r.GET("/pedidos", NewGetAllPedidosController.Execute)
}
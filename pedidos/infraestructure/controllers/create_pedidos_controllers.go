package controllers

import (
	"API_GO/pedidos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePedidosController struct {
	createUseCase   *application.CreatePedidos
	decreaseUseCase *application.DecreaseStock
}

func NewCreatePedidosController(
	createUseCase *application.CreatePedidos,
	decreaseUseCase *application.DecreaseStock,
) *CreatePedidosController {
	return &CreatePedidosController{
		createUseCase:   createUseCase,
		decreaseUseCase: decreaseUseCase,
	}
}

func (c *CreatePedidosController) Execute(ctx *gin.Context) {
	var input struct {
		Estatus           string `json:"estatus"`
		Dessert_id        int    `json:"dessert_id"`
		User_id           int    `json:"user_id"`
		Cantidad_producto int    `json:"cantidad_producto"`
		Total             int    `json:"total"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}
	if err := c.decreaseUseCase.Execute(input.Dessert_id, input.Cantidad_producto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "No se pudo reducir el inventario",
			"detalle": err.Error(),
		})
		return
	}

	if err := c.createUseCase.Execute(input.Dessert_id, input.User_id, input.Cantidad_producto, input.Estatus, input.Total); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el pedido"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Pedido creado correctamente y stock actualizado."})
}

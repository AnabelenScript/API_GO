package controllers

import (
	"API_GO/pedidos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePedidosController struct {
	useCase *application.CreatePedidos
}

func NewCreatePedidosController(useCase *application.CreatePedidos) *CreatePedidosController {
	return &CreatePedidosController{useCase: useCase}
}

func (c *CreatePedidosController) Execute(ctx *gin.Context) {
	var input struct {
		Estatus           string `json:"estatus"`
		Dessert_id        int    `json:"dessert_id"`
		User_id           int    `json:"user_id"`
		Cantidad_producto int    `json:"cantidad_producto"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}
	if err := c.useCase.Execute(input.Dessert_id, input.User_id, input.Cantidad_producto, input.Estatus); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el postre"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Postre creado creado :)"})
}

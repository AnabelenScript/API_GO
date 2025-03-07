package controllers

import (
	"API_GO/pedidos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePedidosController struct {
	useCase *application.UpdatePedidos
}

func NewUpdateUserController(useCase *application.UpdatePedidos) *UpdatePedidosController {
	return &UpdatePedidosController{useCase: useCase}
}

func (c *UpdatePedidosController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var input struct {
		Dessert_id  int `json:"dessert_id"`
		User_id int `json:"user_id"`
		Cantidad_producto int `json:"cantidad_producto"`
		Estatus string `json:"estatus"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}
	if err := c.useCase.Execute(uint(id), input.Dessert_id, input.User_id, input.Cantidad_producto, input.Estatus); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el postre"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Postre actualizado correctamente"})
}

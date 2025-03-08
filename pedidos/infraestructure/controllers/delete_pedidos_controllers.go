package controllers

import (
	"API_GO/pedidos/application"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type DeletePedidoController struct {
	useCase *application.DeletePedido
}

func NewDeleteUserController(useCase *application.DeletePedido) *DeletePedidoController {
	return &DeletePedidoController{useCase: useCase}
}

func (uc *DeletePedidoController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	pedidoID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := uc.useCase.Execute(uint(pedidoID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Postre eliminado correctamente"})
}

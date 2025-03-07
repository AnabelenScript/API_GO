package controllers

import (
	"API_GO/pedidos/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetAllPedidosController struct {
	useCase *application.GetAllPedidos
}

func NewGetAllDessertController(useCase *application.GetAllPedidos) *GetAllPedidosController {
	return &GetAllPedidosController{useCase: useCase}
}

func (uc *GetAllPedidosController) Execute(c *gin.Context) {
	dessert, err := uc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los postres"})
		return
	}

	c.JSON(http.StatusOK, dessert)
}

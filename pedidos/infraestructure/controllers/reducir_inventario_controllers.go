// controllers/decrease_stock_controller.go
package controllers

import (
	"API_GO/pedidos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DecreaseStockController struct {
	useCase *application.DecreaseStock
}

func NewDecreaseStockController(useCase *application.DecreaseStock) *DecreaseStockController {
	return &DecreaseStockController{useCase: useCase}
}

func (c *DecreaseStockController) Execute(ctx *gin.Context) {
	var input struct {
		DessertID int `json:"dessert_id"`
		Quantity  int `json:"cantidad"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	if err := c.useCase.Execute(input.DessertID, input.Quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Stock reducido correctamente"})
}

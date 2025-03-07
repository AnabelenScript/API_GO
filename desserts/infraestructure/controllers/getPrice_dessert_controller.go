package controllers

import (
	"API_GO/desserts/application"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetPriceDessertController struct {
	useCase *application.GetPriceDessert
}

func NewGetPriceDessertController(useCase *application.GetPriceDessert) *GetPriceDessertController {
	return &GetPriceDessertController{useCase: useCase}
}

func (uc *GetPriceDessertController) Execute(c *gin.Context) {
	priceParam := c.Param("price")
	dessertPrice, err := strconv.ParseUint(priceParam, 10, 64)
	dessert, err := uc.useCase.Execute(uint(dessertPrice))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los postres con ese precio"})
		return
	}

	c.JSON(http.StatusOK, dessert)
}

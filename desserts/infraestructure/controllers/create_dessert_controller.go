package controllers

import (
	"API_GO/desserts/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateDessertController struct {
	useCase *application.CreateDessert
}

func NewCreateDessertController(useCase *application.CreateDessert) *CreateDessertController {
	return &CreateDessertController{useCase: useCase}
}

func (c *CreateDessertController) Execute(ctx *gin.Context) {
	var input struct {
		Name  string `json:"name"`
		Flavor string `json:"flavor"`
		Price int `json:"price"`
		Quantity int `json:"quantity"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}
	if err := c.useCase.Execute(input.Name, input.Flavor, input.Price, input.Quantity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Postre creado creado :)"})
}

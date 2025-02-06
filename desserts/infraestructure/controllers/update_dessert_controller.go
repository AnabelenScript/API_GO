package controllers

import (
	"API_GO/desserts/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateDessertController struct {
	useCase *application.UpdateDessert
}

func NewUpdateUserController(useCase *application.UpdateDessert) *UpdateDessertController {
	return &UpdateDessertController{useCase: useCase}
}

func (c *UpdateDessertController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
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
	if err := c.useCase.Execute(uint(id), input.Name, input.Flavor, input.Price, input.Quantity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el postre"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Postre actualizado correctamente"})
}

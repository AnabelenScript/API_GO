package controllers

import (
	"API_GO/desserts/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteDessertController struct {
	useCase *application.DeleteDessert
}

func NewDeleteUserController(useCase *application.DeleteDessert) *DeleteDessertController {
	return &DeleteDessertController{useCase: useCase}
}

func (uc *DeleteDessertController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	dessertID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := uc.useCase.Execute(uint(dessertID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Postre eliminado correctamente"})
}

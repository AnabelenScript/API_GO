package controllers

import (
	"API_GO/desserts/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetAllDessertController struct {
	useCase *application.GetAllDesserts
}

func NewGetAllDessertController(useCase *application.GetAllDesserts) *GetAllDessertController {
	return &GetAllDessertController{useCase: useCase}
}

func (uc *GetAllDessertController) Execute(c *gin.Context) {
	dessert, err := uc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los postres"})
		return
	}

	c.JSON(http.StatusOK, dessert)
}

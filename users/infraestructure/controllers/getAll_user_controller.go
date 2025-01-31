package controllers

import (
	"API_GO/users/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetAllUserController struct {
	useCase *application.GetAllUsers
}

func NewGetAllUserController(useCase *application.GetAllUsers) *GetAllUserController {
	return &GetAllUserController{useCase: useCase}
}

func (uc *GetAllUserController) Execute(c *gin.Context) {
	users, err := uc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	c.JSON(http.StatusOK, users)
}

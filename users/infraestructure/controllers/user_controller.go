package controllers

import (
	"API_GO/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	useCase *application.CreateUser
}

func NewCreateUserController(useCase *application.CreateUser) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

func (c *CreateUserController) Execute(ctx *gin.Context) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}
	if err := c.useCase.Execute(input.Name, input.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario creado :)"})
}

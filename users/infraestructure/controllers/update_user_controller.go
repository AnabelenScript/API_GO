package controllers

import (
	"API_GO/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateUserController struct {
	useCase *application.UpdateUser
}

func NewUpdateUserController(useCase *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{useCase: useCase}
}

func (c *UpdateUserController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}
	if err := c.useCase.Execute(uint(id), input.Name, input.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}

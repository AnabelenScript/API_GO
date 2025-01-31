package controllers

import (
	"API_GO/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	useCase *application.DeleteUser
}

func NewDeleteUserController(useCase *application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{useCase: useCase}
}

func (uc *DeleteUserController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := uc.useCase.Execute(uint(userID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}

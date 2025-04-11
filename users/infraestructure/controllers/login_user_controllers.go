package controllers

import (
	"API_GO/users/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	useCase *application.LoginUser
}

func NewLoginUserController(useCase *application.LoginUser) *LoginUserController {
	return &LoginUserController{useCase: useCase}
}
func (c *LoginUserController) Execute(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}
	user, err := c.useCase.Execute(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"user": gin.H{
			"id":        user.ID,
			"name":      user.Name,
			"email":     user.Email,
			"user_type": user.UserType,
		},
	})
}

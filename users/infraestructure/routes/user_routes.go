package routes

import (
    "API_GO/users/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupUsersRoutes(r *gin.Engine, NewCreateUserController *controllers.CreateUserController) {
    r.POST("/users", NewCreateUserController.Execute)
}

package routes

import (
    "API_GO/users/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupUsersRoutes(
	r *gin.Engine,
	NewCreateUserController *controllers.CreateUserController, 
	NewUpdateUserController *controllers.UpdateUserController, 
	NewDeleteUserController *controllers.DeleteUserController, 
	NewGetAllUserController *controllers.GetAllUserController) {
    r.POST("/users", NewCreateUserController.Execute)
	r.PUT("/users/:id", NewUpdateUserController.Execute)
	r.DELETE("users/:id", NewDeleteUserController.Execute)
	r.GET("/users", NewGetAllUserController.Execute)
}

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
	NewGetAllUserController *controllers.GetAllUserController,
	NewGetLastUserController *controllers.GetLastUserController) {
    r.POST("/users", NewCreateUserController.Execute)
	r.PUT("/users/:id", NewUpdateUserController.Execute)
	r.DELETE("users/:id", NewDeleteUserController.Execute)
	r.GET("/users", NewGetAllUserController.Execute)
	r.GET("/users/getLast", NewGetLastUserController.Execute)

}

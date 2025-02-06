package controllers

import (
	"net/http"
	"strconv"
	"time"

	"API_GO/users/application"
	"github.com/gin-gonic/gin"
)

type GetLastUserController struct {
	Service application.GetLastUser
}

func NewGetLastUserController(uc application.GetLastUser) *GetLastUserController {
	return &GetLastUserController{Service: uc}
}

func (controller *GetLastUserController) Execute(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Encoding", "chunked")

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastUserID int

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{
				"Error": "Timeout: No se encontraron actualizaciones",
			})
			return

		case <-ticker.C:
			result, err := controller.Service.GetLastUserProcess()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Error al obtener el último postre registrado",
				})
				return
			}

			if result.ID > 0 && result.ID != lastUserID {
				lastUserID = result.ID

				idString := strconv.Itoa(result.ID)

				payload := map[string]string{
					"Message":       "Último usuario agregado",
					"User_id":    idString,
					"User_name":  result.Name,
					"USer_flavor": result.Email,
				}

				c.JSON(http.StatusOK, payload)
				return
			}
		}
	}
}

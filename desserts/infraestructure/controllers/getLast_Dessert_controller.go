package controllers

import (
	"net/http"
	"strconv"
	"time"
	"log"
	"API_GO/desserts/application"
	"github.com/gin-gonic/gin"
)

type GetLastDessertController struct {
	Service application.GetLastDesserts
}

func NewGetLastDessertController(uc application.GetLastDesserts) *GetLastDessertController {
	return &GetLastDessertController{Service: uc}
}

func (controller *GetLastDessertController) Execute(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Encoding", "chunked")

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastDessertID int

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{
				"Error": "Timeout: No se encontraron actualizaciones",
			})
			return

		case <-ticker.C:
			log.Print("aaaaa")
			result, err := controller.Service.GetLastDessertProcess()
			
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Error al obtener el último postre registrado",
				})
				return
			}

			if result.Id > 0 && result.Id != lastDessertID {
				log.Println("Cantidad de postres actuales:")
				lastDessertID = result.Id

				idString := strconv.Itoa(result.Id)
				priceString := strconv.Itoa(result.Price)
				quantityString := strconv.Itoa(result.Quantity)

				payload := map[string]string{
					"Message":       "Último postre agregado",
					"Dessert_id":    idString,
					"Dessert_name":  result.Name,
					"Dessert_flavor": result.Flavor,
					"Dessert_price": priceString,
					"Dessert_quantity": quantityString,
				}

				c.JSON(http.StatusOK, payload)
				return
			}
		}
	}
}

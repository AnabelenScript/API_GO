package controllers

import (
	"API_GO/desserts/application"
	"API_GO/desserts/domain/entities" 
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type GetDeletedDessertController struct {
	Service application.GetDeletedDessert
}

func NewGetDeletedDessertController(service application.GetDeletedDessert) *GetDeletedDessertController {
	return &GetDeletedDessertController{Service: service}
}

func (controller *GetDeletedDessertController) Execute(c *gin.Context) {
    c.Header("Content-Type", "text/event-stream")
    c.Header("Cache-Control", "no-cache")
    c.Header("Connection", "keep-alive")

    log.Println("Inicio de la verificación de postres eliminados")

    timeout := time.After(100 * time.Second)
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    var previousDesserts []*entities.Dessert

    for {
        select {
        case <-timeout:
            c.SSEvent("error", gin.H{"message": "Timeout: No se encontraron eliminaciones"})
            return

        case <-ticker.C:
            result, err := controller.Service.GetAllDessertsForDeleteProcess()
            if err != nil {
                c.SSEvent("error", gin.H{"message": "Error al obtener los postres"})
                return
            }

            if len(previousDesserts) == 0 {
                previousDesserts = result
                c.SSEvent("message", gin.H{"Postres": result})
                c.Writer.Flush()
                continue
            }

            if len(result) < len(previousDesserts) {
                c.SSEvent("deleted", gin.H{"Se eliminó un postre; Postres actuales": result})
                c.Writer.Flush()
                previousDesserts = result
            }
        }
    }
}

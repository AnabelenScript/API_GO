package rabbitmq

import (
	"API_GO/pedidos/domain/entities"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type PedidoMensaje struct {
	PedidoID         int    `json:"pedido_id"`
	DessertID        int    `json:"dessert_id"`
	UserID           int    `json:"user_id"`
	CantidadProducto int    `json:"cantidad_producto"`
	Estatus          string `json:"estatus"`
}

func SendPedidoToRabbitMQ(pedido *entities.Pedidos) error {

	conn, err := amqp.Dial("amqp://anita:123456789@52.86.221.36:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
		return err
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(
		"pedidos",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
		return err
	}
	pedidoMensaje := PedidoMensaje{
		PedidoID:         pedido.Pedido_id,
		DessertID:        pedido.Dessert_id,
		UserID:           pedido.User_id,
		CantidadProducto: pedido.Cantidad_producto,
		Estatus:          pedido.Estatus,
	}

	msg, err := json.Marshal(pedidoMensaje)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
		return err
	}

	err = ch.Publish(
		"",
		"pedidos",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
		return err
	}

	log.Println("Mensaje de pedido enviado correctamente.")
	return nil
}

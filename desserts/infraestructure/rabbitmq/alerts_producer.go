package rabbitmq_producer

import (
	"API_GO/desserts/domain/entities"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type PedidoMensaje struct {
	DessertID        int    `json:"dessert_id"`
	Quantity         int    `json: "quantity"`
	Name             string `json:"name"`
	Total            int    `json:"total"`
}

type RabbitMQProducer struct {
	Channel *amqp.Channel
}

func NewRabbitMQProducer(ch *amqp.Channel) *RabbitMQProducer {
	return &RabbitMQProducer{Channel: ch}
}

func (p *RabbitMQProducer) SendPedidoToRabbitMQ(dessert *entities.Dessert) error {
	_, err := p.Channel.QueueDeclare(
		"desserts_alert",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Error al declarar la cola:", err)
		return err
	}
	dessertMensaje := desserMensaje{
		PedidoID:         pedido.Pedido_id,
		DessertID:        pedido.Dessert_id,
		UserID:           pedido.User_id,
		CantidadProducto: pedido.Cantidad_producto,
		Estatus:          pedido.Estatus,
		Total:            pedido.Total,
	}

	msg, err := json.Marshal(pedidoMensaje)
	if err != nil {
		log.Println("Error al convertir a JSON:", err)
		return err
	}
	err = p.Channel.Publish(
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
		log.Println("Error al enviar el pedido:", err)
		return err
	}

	log.Println("Pedido enviado correctamente a RabbitMQ.")
	return nil
}

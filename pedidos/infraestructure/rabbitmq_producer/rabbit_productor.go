package rabbitmq_producer

import (
	"API_GO/pedidos/domain/entities"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// PedidoMensaje representa la estructura del mensaje que se enviará a RabbitMQ
type PedidoMensaje struct {
	PedidoID         int    `json:"pedido_id"`
	DessertID        int    `json:"dessert_id"`
	UserID           int    `json:"user_id"`
	CantidadProducto int    `json:"cantidad_producto"`
	Estatus          string `json:"estatus"`
}

// NewRabbitMQProducer retorna una instancia lista para enviar mensajes
type RabbitMQProducer struct {
	Channel *amqp.Channel
}

// NewRabbitMQProducer inicializa el productor con un canal de RabbitMQ
func NewRabbitMQProducer(ch *amqp.Channel) *RabbitMQProducer {
	return &RabbitMQProducer{Channel: ch}
}

// SendPedidoToRabbitMQ envía un pedido a la cola de RabbitMQ
func (p *RabbitMQProducer) SendPedidoToRabbitMQ(pedido *entities.Pedidos) error {
	// Declarar la cola (si no existe)
	_, err := p.Channel.QueueDeclare(
		"pedidos",
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

	// Convertir el pedido a JSON
	pedidoMensaje := PedidoMensaje{
		PedidoID:         pedido.Pedido_id,
		DessertID:        pedido.Dessert_id,
		UserID:           pedido.User_id,
		CantidadProducto: pedido.Cantidad_producto,
		Estatus:          pedido.Estatus,
	}

	msg, err := json.Marshal(pedidoMensaje)
	if err != nil {
		log.Println("Error al convertir a JSON:", err)
		return err
	}

	// Publicar el mensaje en la cola
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

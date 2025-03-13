package helpers

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://anita:123456789@52.86.221.36:5672/")
	if err != nil {
		log.Println("Error al conectar con RabbitMQ:", err)
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error al abrir un canal de RabbitMQ:", err)
		conn.Close()
		return nil, nil, err
	}
	return conn, ch, nil
}

package db

import (
	"database/sql"
	"API_GO/pedidos/domain"
	"API_GO/pedidos/domain/entities"
	"log"
	"errors"
	"API_GO/pedidos/infraestructure/rabbitmq_producer"
	"github.com/streadway/amqp"
)

type MySQLPedidosRepository struct {
	DB *sql.DB
	RabbitMQProducer *rabbitmq_producer.RabbitMQProducer
}

func NewMySQLPedidosRepository(db *sql.DB, ch *amqp.Channel) domain.PedidosRepository {
	return &MySQLPedidosRepository{
		RabbitMQProducer: rabbitmq_producer.NewRabbitMQProducer(ch),
		DB: db}	
}

func (r *MySQLPedidosRepository) Save(pedido *entities.Pedidos) error {
	query := "INSERT INTO pedidos (dessert_id, user_id, cantidad_producto, estatus) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, pedido.Dessert_id, pedido.User_id, pedido.Cantidad_producto, pedido.Estatus)
	if err != nil {
		log.Printf("Error al agregar el pedido: %v", err)
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID autogenerado: %v", err)
		return err
	}
	pedido.Pedido_id = int(lastInsertID)
	err = r.RabbitMQProducer.SendPedidoToRabbitMQ(pedido)
	if err != nil {
		log.Printf("Error al enviar el pedido a RabbitMQ: %v", err)
		return err
	}
	log.Println("Pedido guardado y enviado a RabbitMQ.")
	return nil
}


func (r *MySQLPedidosRepository) FindByID(pedido_id uint) (*entities.Pedidos, error) {
	query := "SELECT pedido_id, dessert_id, user_id, cantidad_producto, estatus FROM pedidos WHERE pedido_id = ?"
	row := r.DB.QueryRow(query, pedido_id)

	var pedidos entities.Pedidos
	err := row.Scan(&pedidos.Pedido_id, &pedidos.Dessert_id, &pedidos.User_id, &pedidos.Cantidad_producto, &pedidos.Estatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Postre no encontrado :C")
		}
		log.Printf("Error al buscar postre: %v", err)
		return nil, err
	}
	return &pedidos, nil
}

func (r *MySQLPedidosRepository) Update(pedido *entities.Pedidos) error {
	query := "UPDATE pedidos SET dessert_id = ?, user_id = ?, cantidad_producto = ?, estatus = ? WHERE pedido_id = ?"
	result, err := r.DB.Exec(query, pedido.Dessert_id, pedido.User_id, pedido.Cantidad_producto, pedido.Estatus)
	if err != nil {
		log.Printf("Error al actualizar el postre: %v", err)
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no se encontró el postre para actualizar :(")
	}

	return nil
}

func (r *MySQLPedidosRepository) Delete(pedidoID uint) error {
	query := "DELETE FROM pedidos WHERE pedido_id = ?"
	result, err := r.DB.Exec(query, pedidoID)
	if err != nil {
		log.Printf("Error al eliminar el postre: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no se encontró el postre para eliminar")
	}

	return nil
}

func (r *MySQLPedidosRepository) GetAll() ([]*entities.Pedidos, error) {
	query := "SELECT pedido_id, dessert_id, user_id, cantidad_productos, estatus FROM pedidos"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Error al obtener todos los postres: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pedidos []*entities.Pedidos
	for rows.Next() {
		pedido := &entities.Pedidos{}
		if err := rows.Scan(&pedido.Pedido_id, &pedido.Dessert_id, &pedido.User_id, &pedido.Cantidad_producto, &pedido.Estatus); err != nil {
			log.Printf("Error al escanear el postre: %v", err)
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	return pedidos, nil
}

package db


import (
	"database/sql"
	"API_GO/pedidos/domain"
	"API_GO/pedidos/domain/entities"
	"log"
	"errors"
)

type MySQLPedidosRepository struct {
	DB *sql.DB
}

func NewMySQLPedidosRepository(db *sql.DB) domain.PedidosRepository {
	return &MySQLPedidosRepository{DB: db}
}


func (r *MySQLPedidosRepository) Save(pedidos *entities.Pedidos) error {
	query := "INSERT INTO pedidos (dessert_id, user_id, cantidad_producto, estatus) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, pedidos.Dessert_id, pedidos.User_id, pedidos.Cantidad_producto, pedidos.Estatus)
	if err != nil {
		log.Printf("Error al agrgear el postre: %v", err)
	}
	return err
}

func (r *MySQLPedidosRepository) FindByID(id uint) (*entities.Pedidos, error) {
	query := "SELECT ID, dessert_id, user_id, cantidad_producto, estatus FROM pedidos WHERE ID = ?"
	row := r.DB.QueryRow(query, id)

	var pedidos entities.Pedidos
	err := row.Scan(&pedidos.Id, &pedidos.Dessert_id, &pedidos.User_id, &pedidos.Cantidad_producto, &pedidos.Estatus)
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
	query := "UPDATE pedidos SET dessert_id = ?, user_id = ?, cantidad_producto = ?, estatus = ? WHERE ID = ?"
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
	query := "DELETE FROM pedidos WHERE ID = ?"
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
	query := "SELECT ID, dessert_id, user_id, cantidad_productos, estatus FROM pedidos"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Error al obtener todos los postres: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pedidos []*entities.Pedidos
	for rows.Next() {
		pedido := &entities.Pedidos{}
		if err := rows.Scan(&pedido.Id, &pedido.Dessert_id, &pedido.User_id, &pedido.Cantidad_producto, &pedido.Estatus); err != nil {
			log.Printf("Error al escanear el postre: %v", err)
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	return pedidos, nil
}

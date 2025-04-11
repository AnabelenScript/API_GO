package domain

import "API_GO/pedidos/domain/entities"

type PedidosRepository interface {
	Save(pedidos *entities.Pedidos) error
	FindByID(pedido_id uint) (*entities.Pedidos, error)
	Update(pedidos *entities.Pedidos) error
	Delete(pedido_id uint) error
	GetAll() ([]*entities.Pedidos, error)
	DecreaseDessertStock(dessertID int, cantidad int) error
}
package domain

import "API_GO/pedidos/domain/entities"

type PedidosRepository interface {
	Save(pedidos *entities.Pedidos) error
	FindByID(id uint) (*entities.Pedidos, error)
	Update(pedidos *entities.Pedidos) error
	Delete(id uint) error
	GetAll() ([]*entities.Pedidos, error)
}
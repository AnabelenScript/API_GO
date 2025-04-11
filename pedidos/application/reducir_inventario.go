package application

import "API_GO/pedidos/domain"

type DecreaseStock struct {
	inventario domain.PedidosRepository
}

func NewDecreaseStock(inventory domain.PedidosRepository) *DecreaseStock {
	return &DecreaseStock{inventario: inventory}
}

func (ds *DecreaseStock) Execute(dessertID int, cantidad int) error {
	return ds.inventario.DecreaseDessertStock(dessertID, cantidad)
}

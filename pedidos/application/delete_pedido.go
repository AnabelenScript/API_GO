package application

import (
	"API_GO/pedidos/domain"
	"errors"
)

type DeletePedido struct {
	Repo domain.PedidosRepository
	
}

func NewDeletePedido(repo domain.PedidosRepository) *DeletePedido {
	return &DeletePedido{Repo: repo}
}

func (uc *DeletePedido) Execute(pedido_id uint) error {
	pedido, err := uc.Repo.FindByID(pedido_id)
	if err != nil {
		return errors.New("Postre no encontrado")
	}
	return uc.Repo.Delete(uint(pedido.Pedido_id))
}

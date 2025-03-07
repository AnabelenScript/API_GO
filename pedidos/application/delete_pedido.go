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

func (uc *DeletePedido) Execute(id uint) error {
	pedido, err := uc.Repo.FindByID(id)
	if err != nil {
		return errors.New("Postre no encontrado")
	}
	return uc.Repo.Delete(uint(pedido.Id))
}

package application

import (
	"API_GO/pedidos/domain"
)

type UpdatePedidos struct {
	Repo domain.PedidosRepository
}

func NewUpdatePedidos(repo domain.PedidosRepository) *UpdatePedidos {
	return &UpdatePedidos{Repo: repo}
}

func (uc *UpdatePedidos) Execute(pedido_id uint, dessert_id int, user_id int, cantidad_producto int, estatus string) error {
	pedido, err := uc.Repo.FindByID(pedido_id)
	if err != nil {
		return err
	}
	pedido.Dessert_id = dessert_id
	pedido.User_id = user_id
	pedido.Cantidad_producto = cantidad_producto
	pedido.Estatus = estatus
	return uc.Repo.Update(pedido)
}

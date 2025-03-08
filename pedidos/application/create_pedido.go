package application

import (
	"API_GO/pedidos/domain"
	"API_GO/pedidos/domain/entities"
)

type CreatePedidos struct {
	Repo domain.PedidosRepository
}

func NewCreatePedidos(repo domain.PedidosRepository) *CreatePedidos {
	return &CreatePedidos{Repo: repo}
}

func (uc *CreatePedidos) Execute(dessert_id int, user_id int, cantidad_producto int, estatus string) error {
	pedidos := &entities.Pedidos{Dessert_id: dessert_id, User_id: user_id, Cantidad_producto: cantidad_producto, Estatus: estatus}
	return uc.Repo.Save(pedidos)
}

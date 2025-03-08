package application

import (
	"API_GO/pedidos/domain"
	"API_GO/pedidos/domain/entities"
)

type GetAllPedidos struct {
	Repo domain.PedidosRepository
}

func NewGetAllPedidos(repo domain.PedidosRepository) *GetAllPedidos {
	return &GetAllPedidos{Repo: repo}
}

func (uc *GetAllPedidos) Execute() ([]*entities.Pedidos, error) {
	return uc.Repo.GetAll()
}

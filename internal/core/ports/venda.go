package ports

import (
	"context"
	"soat_substitutiva/internal/core/domain"
)

type VendaService interface {
	BuscarVendas(ctx context.Context) ([]domain.Venda, error)
	CriarVenda(ctx context.Context, venda domain.Venda) error
}

type VendaRepository interface {
	BuscarVendas(ctx context.Context) ([]domain.Venda, error)
	CriarVenda(ctx context.Context, venda domain.Venda) error
}

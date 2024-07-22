package ports

import (
	"context"
	"soat_substitutiva/internal/core/domain"

	"github.com/labstack/echo/v4"
)

type VendaHandler interface {
	BuscarVendas(echo.Context) error
	CriarVenda(echo.Context) error
}
type VendaService interface {
	BuscarVendas(ctx context.Context) ([]domain.Venda, error)
	CriarVenda(ctx context.Context, venda domain.Venda) error
}

type VendaRepository interface {
	BuscarVendas(ctx context.Context) ([]domain.Venda, error)
	CriarVenda(ctx context.Context, venda domain.Venda) error
}

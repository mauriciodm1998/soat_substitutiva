package ports

import (
	"context"
	"soat_substitutiva/internal/core/domain"

	"github.com/labstack/echo/v4"
)

type ClienteHandler interface {
	CadastrarCliente(echo.Context) error
}

type ClienteService interface {
	CadastrarCliente(ctx context.Context, customer domain.Cliente) (string, error)
}

type ClienteRepository interface {
	CadastrarCliente(ctx context.Context, customer domain.Cliente) error
}

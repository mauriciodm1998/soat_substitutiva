package ports

import (
	"context"
	"soat_substitutiva/internal/core/domain"
)

type ClienteService interface {
	CadastrarCliente(ctx context.Context, customer domain.Cliente) (string, error)
}

type ClienteRepository interface {
	CadastrarCliente(ctx context.Context, customer domain.Cliente) error
}

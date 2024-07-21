package ports

import (
	"context"
	"soat_substitutiva/internal/core/domain"
)

type VeiculoService interface {
	RegistrarVeiculo(context.Context, domain.Veiculo) (string, error)
	ListarVeiculosAVenda(context.Context) ([]domain.Veiculo, error)
	ListarVeiculosVendidos(context.Context) ([]domain.Veiculo, error)
	EditarVeiculo(context.Context, domain.Veiculo) error
}

type VeiculoRepository interface {
	BuscarPorId(context.Context, string) (*domain.Veiculo, error)
	EditarVeiculo(context.Context, domain.Veiculo) error
	RegistrarVeiculo(ctx context.Context, vehicle domain.Veiculo) error
	ListarVeiculosAVenda(context.Context) ([]domain.Veiculo, error)
	ListarVeiculosVendidos(context.Context) ([]domain.Veiculo, error)
}

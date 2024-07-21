package usecases

import (
	"context"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type veiculoService struct {
	repo ports.VeiculoRepository
}

func NewVeiculoService(repo ports.VeiculoRepository) ports.VeiculoService {
	return &veiculoService{repo: repo}
}

func (s *veiculoService) RegistrarVeiculo(ctx context.Context, veiculo domain.Veiculo) (string, error) {
	veiculo.ID = uuid.New().String()
	veiculo.DataDeCriacao = time.Now()
	veiculo.DataDeModificacao = time.Now()

	err := s.repo.RegistrarVeiculo(ctx, veiculo)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao registrar veiculo")
		return "", err
	}

	return veiculo.ID, nil
}

func (s *veiculoService) ListarVeiculosAVenda(ctx context.Context) ([]domain.Veiculo, error) {
	veiculos, err := s.repo.ListarVeiculosAVenda(ctx)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao buscar os veiculos a venda")
		return nil, err
	}

	return veiculos, nil
}

func (s *veiculoService) ListarVeiculosVendidos(ctx context.Context) ([]domain.Veiculo, error) {
	veiculos, err := s.repo.ListarVeiculosVendidos(ctx)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao buscar os veiculos vendidos")
		return nil, err
	}

	return veiculos, nil
}

func (s *veiculoService) EditarVeiculo(ctx context.Context, veiculo domain.Veiculo) error {
	veic, err := s.repo.BuscarPorId(ctx, veiculo.ID)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao buscar o veiculo")
		return err
	}

	veic.DataDeModificacao = time.Now()

	err = s.repo.EditarVeiculo(ctx, veiculo)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao editar o veiculo")
		return err
	}

	return nil
}

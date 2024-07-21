package usecases

import (
	"context"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type vendaService struct {
	repo              ports.VendaRepository
	veiculoRepository ports.VeiculoRepository
}

func NewVendaService(repo ports.VendaRepository, veiculoRepo ports.VeiculoRepository) ports.VendaService {
	return &vendaService{repo: repo, veiculoRepository: veiculoRepo}
}

func (s *vendaService) CriarVenda(ctx context.Context, venda domain.Venda) error {
	venda.ID = uuid.New().String()
	venda.Data = time.Now()

	err := s.repo.CriarVenda(ctx, venda)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao criar venda")
		return err
	}

	veic, err := s.veiculoRepository.BuscarPorId(ctx, venda.Veiculo.ID)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao buscar pelo veiculo")
		return err
	}

	veic.Disponivel = false
	veic.DataDeModificacao = time.Now()

	err = s.veiculoRepository.EditarVeiculo(ctx, *veic)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao editar o veiculo")
		return err
	}

	return nil
}

func (s *vendaService) BuscarVendas(ctx context.Context) ([]domain.Venda, error) {
	vendas, err := s.repo.BuscarVendas(ctx)
	if err != nil {
		logrus.WithError(err).Error("um erro ocorreu ao buscar pelas vendas")
		return nil, err
	}

	return vendas, nil
}

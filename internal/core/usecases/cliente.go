package usecases

import (
	"context"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type clienteService struct {
	repo ports.ClienteRepository
}

func NewClienteService(repo ports.ClienteRepository) ports.ClienteService {
	return &clienteService{repo: repo}
}

func (s *clienteService) CadastrarCliente(ctx context.Context, cliente domain.Cliente) (string, error) {
	cliente.ID = uuid.New().String()
	cliente.DataDeCriacao = time.Now()

	err := s.repo.CadastrarCliente(ctx, cliente)
	if err != nil {
		logrus.WithError(err).Error("erro ao cadastrar usuario no banco")
		return "", err
	}

	return cliente.ID, nil
}

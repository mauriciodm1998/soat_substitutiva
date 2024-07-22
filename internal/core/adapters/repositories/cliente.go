package repositories

import (
	"context"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type clienteRepository struct {
	db *pgxpool.Pool
}

func NewClienteRepository(db *pgxpool.Pool) ports.ClienteRepository {
	return &clienteRepository{db: db}
}

func (r *clienteRepository) CadastrarCliente(ctx context.Context, cliente domain.Cliente) error {
	sqlStatement := "INSERT INTO \"Cliente\" (id, nome, documento, data_de_criacao) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(ctx, sqlStatement, cliente.ID, cliente.Nome, cliente.Documento, cliente.DataDeCriacao)
	if err != nil {
		return err
	}

	return nil
}

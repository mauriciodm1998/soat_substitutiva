package repositories

import (
	"context"
	"fmt"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type veiculoRepository struct {
	db *pgxpool.Pool
}

func NewVeiculoRepository(db *pgxpool.Pool) ports.VeiculoRepository {
	return &veiculoRepository{db: db}
}

func (r *veiculoRepository) RegistrarVeiculo(ctx context.Context, veiculo domain.Veiculo) error {
	sqlStatement := "INSERT INTO \"Veiculo\" (id, marca, modelo, ano, cor, preco, disponivel) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.db.Exec(ctx, sqlStatement, veiculo.ID, veiculo.Marca, veiculo.Modelo, veiculo.Ano, veiculo.Cor, veiculo.Preco, veiculo.Disponivel)
	if err != nil {
		return err
	}

	return nil
}

func (r *veiculoRepository) ListarVeiculosAVenda(ctx context.Context) ([]domain.Veiculo, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM \"Veiculo\" WHERE disponivel = true ORDER BY preco ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var veiculos []domain.Veiculo

	for rows.Next() {
		var veiculo domain.Veiculo

		if err = rows.Scan(
			&veiculo.ID,
			&veiculo.Marca,
			&veiculo.Modelo,
			&veiculo.Ano,
			&veiculo.Cor,
			&veiculo.Preco,
			&veiculo.Disponivel,
		); err != nil {
			return nil, err
		}

		veiculos = append(veiculos, veiculo)
	}

	return veiculos, nil
}

func (r *veiculoRepository) ListarVeiculosVendidos(ctx context.Context) ([]domain.Veiculo, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM \"Veiculo\" WHERE disponivel = false ORDER BY preco ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var veiculos []domain.Veiculo

	for rows.Next() {
		var veiculo domain.Veiculo

		if err = rows.Scan(
			&veiculo.ID,
			&veiculo.Marca,
			&veiculo.Modelo,
			&veiculo.Ano,
			&veiculo.Cor,
			&veiculo.Preco,
			&veiculo.Disponivel,
		); err != nil {
			return nil, err
		}

		veiculos = append(veiculos, veiculo)
	}

	return veiculos, nil
}

func (r *veiculoRepository) EditarVeiculo(ctx context.Context, veiculo domain.Veiculo) error {
	sqlStatement := "UPDATE \"Veiculo\" SET marca = $1, modelo = $2, ano = $3, cor = $4, preco = $5, disponivel = $6 WHERE id = $7"

	_, err := r.db.Exec(ctx, sqlStatement,
		veiculo.Marca, veiculo.Modelo, veiculo.Ano, veiculo.Cor, veiculo.Preco, veiculo.Disponivel, veiculo.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar o veículo: %w", err)
	}

	return nil
}

func (r *veiculoRepository) BuscarPorId(ctx context.Context, id string) (*domain.Veiculo, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Veiculo\" WHERE id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var veiculo domain.Veiculo
	if rows.Next() {
		if err = rows.Scan(
			&veiculo.ID,
			&veiculo.Marca,
			&veiculo.Modelo,
			&veiculo.Ano,
			&veiculo.Cor,
			&veiculo.Preco,
			&veiculo.Disponivel,
		); err != nil {
			return nil, err
		}
		return &veiculo, nil
	}

	return nil, err
}

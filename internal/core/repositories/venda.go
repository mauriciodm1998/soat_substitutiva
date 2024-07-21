package repositories

import (
	"context"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type vendaRepository struct {
	db *pgxpool.Pool
}

func NewVendaRepository(db *pgxpool.Pool) ports.VendaRepository {
	return &vendaRepository{db: db}
}

func (r *vendaRepository) CriarVenda(ctx context.Context, venda domain.Venda) error {
	sqlStatement := "INSERT INTO \"Venda\" (id, data, veiculo_id, cliente_id, tipo_de_pagamento) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.Exec(ctx, sqlStatement, venda.ID, venda.Data, venda.Veiculo.ID, venda.Cliente.ID, venda.TipoDePagamento)
	if err != nil {
		return err
	}

	return nil
}

func (r *vendaRepository) BuscarVendas(ctx context.Context) ([]domain.Venda, error) {
	rows, err := r.db.Query(ctx, "SELECT v.id, v.marca, v.modelo, v.ano, v.cor, v.preco, v.disponivel, c.id, c.nome, c.documento, c.data_de_criacao, vd.id, vd.data, vd.tipo_de_pagamento FROM \"Venda\" vd INNER JOIN \"Veiculo\" v ON vd.veiculo_id = v.id INNER JOIN \"Cliente\" c ON vd.cliente_id = c.id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var vendas []domain.Venda

	for rows.Next() {
		var venda domain.Venda

		if err = rows.Scan(
			&venda.Veiculo.ID,
			&venda.Veiculo.Marca,
			&venda.Veiculo.Modelo,
			&venda.Veiculo.Ano,
			&venda.Veiculo.Cor,
			&venda.Veiculo.Preco,
			&venda.Veiculo.Disponivel,
			&venda.Cliente.ID,
			&venda.Cliente.Nome,
			&venda.Cliente.Documento,
			&venda.Cliente.DataDeCriacao,
			&venda.ID,
			&venda.Data,
			&venda.TipoDePagamento,
		); err != nil {
			return nil, err
		}

		vendas = append(vendas, venda)
	}

	return vendas, nil
}

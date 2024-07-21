package domain

import "time"

type Veiculo struct {
	ID                string
	Marca             string
	Modelo            string
	Ano               int
	Cor               string
	Preco             float64
	Disponivel        bool
	DataDeCriacao     time.Time
	DataDeModificacao time.Time
}

type Cliente struct {
	ID            string
	Nome          string
	Documento     string
	DataDeCriacao time.Time
}

type Venda struct {
	ID              string
	Data            time.Time
	Veiculo         Veiculo
	Cliente         Cliente
	TipoDePagamento string
}

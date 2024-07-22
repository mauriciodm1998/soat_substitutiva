package rest

type ClienteRequest struct {
	Nome      string `json:"nome"`
	Documento string `json:"documento"`
}

type GenericId struct {
	Id string `json:"id"`
}

type VeiculoRequest struct {
	Marca  string  `json:"marca"`
	Modelo string  `json:"modelo"`
	Ano    int     `json:"ano"`
	Cor    string  `json:"cor"`
	Preco  float64 `json:"preco"`
}

type VendaRequest struct {
	VeiculoId       string `json:"veiculo_id"`
	ClienteId       string `json:"cliente_id"`
	TipoDePagamento string `json:"tipo_de_pagamento"`
}

CREATE TABLE IF NOT EXISTS "Cliente" (
    id VARCHAR(255) PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    documento VARCHAR(50) NOT NULL,
    data_de_criacao TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "Veiculo"(
    id VARCHAR(255) PRIMARY KEY,
    marca VARCHAR(100) NOT NULL,
    modelo VARCHAR(100) NOT NULL,
    ano INT NOT NULL,
    cor VARCHAR(50) NOT NULL,
    preco DECIMAL(10, 2) NOT NULL,
    disponivel BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS "Venda" (
    id VARCHAR(255) PRIMARY KEY,
    data TIMESTAMP NOT NULL,
    veiculo_id VARCHAR(255) NOT NULL,
    cliente_id VARCHAR(255) NOT NULL,
    tipo_de_pagamento VARCHAR(50) NOT NULL,
    FOREIGN KEY (veiculo_id) REFERENCES "Veiculo" (id),
    FOREIGN KEY (cliente_id) REFERENCES "Cliente" (id)
);

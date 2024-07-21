# FIAP - SOAT-POSTECH - Atividade substitutiva

## Descricao

Esse serviço é responsavel pela manipulação de veiculos, clientes e vendas. Ele utiliza arquitetura hexagonal para poder separar melhor as responsabilidades de cada etapa e ser possivel implementar novas funcionalidades de maneira mais facil. Ele sobe um server http e distribui as chamadas entre os 3 dominios diferentes de acordo com a rota: /veiculos, /clientes, /venda.
A linguagem obiqua está sendo respeitada e o que foi discutido no momento do event storming e definições de palavras chave estão sendo devidamente implementadas no código.

## Features

- CadastrarCliente

- RegistrarVeiculo
- ListarVeiculosAVenda
- ListarVeiculosVendidos
- EditarVeiculo

- CriarVenda

## Como rodar localmente

É possivel buildar localmente as imagens para posteriormente utiliza-las no docker-compose, ou utilizar diretamente as imagens do repositório remoto:

Opção 1 para api: $```docker build -t api . -f Dockerfile```
Opção 1 para db: $```docker build -t db . -f Dockerfile.db```

Opção 2 para ambos: $```make run-app```

ou rode manualmente o comando: $```mdocker-compose -f build/db-docker-compose.yaml up -d```

Ambos terão o mesmo resultado.

## Testando a api manualmente

No diretorio api é possivel encontrar uma collection contendo todos os endpoints da aplicação.

## Build + Bake Image

O Dockerfile criado para a aplicação trabalha com duas etapas, a primeira cria um container para que seja possivel executar o build da aplicação e gerar seu respectivo binário, em seguida é executado outro container que pega o binário gerado e o coloca em uma outra imagem menor.
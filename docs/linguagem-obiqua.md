# Documentação da Linguagem Ubíqua

## Introdução
A linguagem ubíqua é um vocabulário comum utilizado por todos os membros da equipe para garantir que todos entendam e se comuniquem de forma consistente sobre os conceitos e processos do domínio de negócio. Esta seção define os principais termos e conceitos utilizados na plataforma de revenda de veículos automotores.

## Termos e Definições

### Veículo
- **Descrição:** Um automóvel disponível para venda ou vendido na plataforma.
- **Atributos:**
  - **Marca:** Fabricante do veículo (ex: Ford, Chevrolet).
  - **Modelo:** Modelo específico do veículo (ex: Fiesta, Onix).
  - **Ano:** Ano de fabricação do veículo.
  - **Cor:** Cor do veículo.
  - **Preco:** Preço de venda do veículo.
  - **Disponivel** Se o veiculo esta disponivel ou nao

### ListarVeiculosAVenda
- **Descrição:** Exibição dos veículos disponíveis para venda.
- **Eventos Relacionados:** veiculos a venda foram listados.

### ListarVeiculosVendidos
- **Descrição:** Exibição dos veículos que já foram vendidos.
- **Eventos Relacionados:** veiculos vendidos foram listados.

### RegistrarVeículo
- **Descrição:** Processo de adicionar um novo veículo à plataforma para que esteja disponível para venda.
- **Eventos Relacionados:** Veículo foi registrado.

### EditarVeículo
- **Descrição:** Processo de atualizar as informações de um veículo já cadastrado na plataforma.
- **Eventos Relacionados:** Veículo foi editado.

### Cliente
- **Descrição:** Um cliente usuario da plataforma.
- **Atributos:**
  - **Nome:** Nome da pessoa que comprou o veículo.
  - **Documento:** CPF da pessoa que comprou o veículo.
  - **Data:** Data de cadastrado.

### CadastrarCliente
- **Descrição:** Processo de cadastrar clientes na plataforma.
- **Eventos Relacionados:** Cliente foi cadastrado.

### Venda
- **Descrição:** Uma venda gerada pela plataforma.
- **Atributos:**
  - **Nome do Cliente:** Nome da pessoa que comprou o veículo.
  - **Documento do Cliente:** CPF da pessoa que comprou o veículo.
  - **Data:** Data em que a venda foi efetuada.
  - **Veiculo** Veiculo que foi comprado

### CriarVenda
- **Descrição:** Processo de criar uma nova venda à plataforma.
- **Eventos Relacionados:** cliente comprou veiculo; venda foi gerada.

## Atores

### Operador
- **Descrição:** Pessoa responsável por cadastrar e editar veiculos na plataforma.

### Cliente
- **Descrição:** Pessoa que compra um veículo através da plataforma.
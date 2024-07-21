package main

import (
	"context"
	"soat_substitutiva/internal/adapters/rest"
	"soat_substitutiva/internal/config"
	"soat_substitutiva/internal/core/ports"
	"soat_substitutiva/internal/core/repositories"
	"soat_substitutiva/internal/core/usecases"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()

	router := echo.New()

	clienteService, vendaService, veiculoService := startDependencies()

	rest.NewClienteHdl(router, clienteService)
	rest.NewVendaHdl(router, vendaService)
	rest.NewVeiculoHdl(router, veiculoService)

	router.Start(":8080")
}

func startDependencies() (ports.ClienteService, ports.VendaService, ports.VeiculoService) {
	db, err := pgxpool.New(context.Background(), config.Get().DB.ConnectionString)
	if err != nil {
		logrus.Fatal(err)
	}

	// cliente
	clienteRepository := repositories.NewClienteRepository(db)
	clienteService := usecases.NewClienteService(clienteRepository)

	//veiculo
	veiculoRepository := repositories.NewVeiculoRepository(db)
	veiculoService := usecases.NewVeiculoService(veiculoRepository)

	//venda
	vendaRepository := repositories.NewVendaRepository(db)
	vendaService := usecases.NewVendaService(vendaRepository, veiculoRepository)

	return clienteService, vendaService, veiculoService
}

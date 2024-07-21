package rest

import (
	"net/http"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type veiculoHdl struct {
	veiculoService ports.VeiculoService
}

func NewVeiculoHdl(router *echo.Echo, veiculoService ports.VeiculoService) {
	handler := &veiculoHdl{
		veiculoService: veiculoService,
	}

	api := router.Group("/veiculos")
	api.POST("/", handler.RegistrarVeiculo)
	api.GET("/", handler.ListarVeiculosAVenda)
	api.GET("/vendidos", handler.ListarVeiculosVendidos)
	api.PUT("/vendidos", handler.ListarVeiculosVendidos)
}

func (h *veiculoHdl) RegistrarVeiculo(c echo.Context) error {
	var request VeiculoRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	id, err := h.veiculoService.RegistrarVeiculo(c.Request().Context(), domain.Veiculo{
		Marca:      request.Marca,
		Modelo:     request.Modelo,
		Ano:        request.Ano,
		Cor:        request.Cor,
		Preco:      request.Preco,
		Disponivel: true,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, GenericId{
		Id: id,
	})
}

func (h *veiculoHdl) ListarVeiculosAVenda(c echo.Context) error {
	veiculos, err := h.veiculoService.ListarVeiculosAVenda(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, veiculos)
}

func (h *veiculoHdl) ListarVeiculosVendidos(c echo.Context) error {
	veiculos, err := h.veiculoService.ListarVeiculosVendidos(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, veiculos)
}

func (h *veiculoHdl) EditarVeiculo(c echo.Context) error {
	var request VeiculoRequest

	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})

	}

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	err = h.veiculoService.EditarVeiculo(c.Request().Context(), domain.Veiculo{
		ID:         id,
		Marca:      request.Marca,
		Modelo:     request.Modelo,
		Ano:        request.Ano,
		Cor:        request.Cor,
		Preco:      request.Preco,
		Disponivel: true,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.NoContent(http.StatusNoContent)
}

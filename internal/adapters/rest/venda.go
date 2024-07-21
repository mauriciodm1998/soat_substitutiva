package rest

import (
	"net/http"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type VendaHdl struct {
	vendaService ports.VendaService
}

func NewVendaHdl(router *echo.Echo, VendaService ports.VendaService) {
	handler := &VendaHdl{
		vendaService: VendaService,
	}

	api := router.Group("/venda")
	api.POST("/", handler.CriarVenda)
	api.GET("/", handler.BuscarVendas)
}

func (h *VendaHdl) CriarVenda(c echo.Context) error {
	var request VendaRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	err = h.vendaService.CriarVenda(c.Request().Context(), domain.Venda{
		Veiculo: domain.Veiculo{
			ID: request.VeiculoId,
		},
		Cliente: domain.Cliente{
			ID: request.ClienteId,
		},
		TipoDePagamento: request.TipoDePagamento,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.NoContent(http.StatusOK)
}

func (h *VendaHdl) BuscarVendas(c echo.Context) error {
	vendas, err := h.vendaService.BuscarVendas(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, vendas)
}

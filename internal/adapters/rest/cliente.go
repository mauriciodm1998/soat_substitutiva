package rest

import (
	"net/http"
	"soat_substitutiva/internal/core/domain"
	"soat_substitutiva/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type clienteHdl struct {
	clienteService ports.ClienteService
}

func NewClienteHdl(router *echo.Echo, clienteService ports.ClienteService) {
	handler := &clienteHdl{
		clienteService: clienteService,
	}

	api := router.Group("/clientes")
	api.POST("/", handler.CadastrarCliente)
}

func (h *clienteHdl) CadastrarCliente(c echo.Context) error {
	var request ClienteRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	id, err := h.clienteService.CadastrarCliente(c.Request().Context(), domain.Cliente{
		Nome:      request.Nome,
		Documento: request.Documento,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, GenericId{
		Id: id,
	})
}

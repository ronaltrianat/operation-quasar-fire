package http

import (
	net_http "net/http"

	"github.com/labstack/echo/v4"

	"github.com/ronaltrianat/operation-quasar-fire/src/core/ports"
)

type http struct {
	topSecretPort ports.TopSecretPort
}

func NewHTTPHandler(service ports.TopSecretPort) *http {
	return &http{
		topSecretPort: service,
	}
}

func (handler *http) TopSecret(c echo.Context) error {
	request := new(TopSecretRequest)
	if err := c.Bind(request); err != nil {
		return c.NoContent(net_http.StatusBadRequest)
	}

	if err := request.Validate(); err != nil {
		return c.NoContent(net_http.StatusBadRequest)
	}

	satellitesData := request.ToDomainData()
	response, err := handler.topSecretPort.TopSecret(satellitesData)

	if err != nil {
		return c.NoContent(net_http.StatusInternalServerError)
	}

	return c.JSON(net_http.StatusOK, response)
}

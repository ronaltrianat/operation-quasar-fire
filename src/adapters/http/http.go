package http

import (
	net_http "net/http"

	"github.com/labstack/echo/v4"

	"github.com/ronaltrianat/operation-quasar-fire/src/core/ports"
)

const (
	namePathParameter           = "name"
	nameQueryParameters         = "satellites"
	minNumberSatellitesRequired = 3
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

func (handler *http) SaveTopSecretSplit(c echo.Context) error {
	name := c.Param(namePathParameter)
	if len(name) == 0 {
		return c.NoContent(net_http.StatusBadRequest)
	}

	request := new(SatelliteRequest)
	if err := c.Bind(request); err != nil {
		return c.NoContent(net_http.StatusBadRequest)
	}

	request.Name = name
	if err := request.Validate(); err != nil {
		return c.NoContent(net_http.StatusBadRequest)
	}

	satellite := request.ToDomainData(name)

	err := handler.topSecretPort.SaveTopSecretSplit(satellite)

	if err != nil {
		return c.NoContent(net_http.StatusInternalServerError)
	}

	return c.NoContent(net_http.StatusOK)
}

func (handler *http) GetTopSecretSplit(c echo.Context) error {
	urlQuery := c.Request().URL.Query()
	satellites := urlQuery[nameQueryParameters]

	if len(satellites) != minNumberSatellitesRequired {
		return c.NoContent(net_http.StatusBadRequest)
	}

	response, err := handler.topSecretPort.RetrieveTopSecretSplit(satellites)

	if err != nil {
		return c.String(net_http.StatusNotFound, err.Error())
	}

	return c.JSON(net_http.StatusOK, response)
}

func (handler *http) Health(c echo.Context) error {
	return c.String(net_http.StatusOK, "OK")
}

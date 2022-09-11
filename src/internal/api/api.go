package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jgcaceres97/go-inventory-api/src/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv          service.Service
	dataValidator *validator.Validate
}

func New(serv service.Service) *API {
	return &API{serv: serv, dataValidator: validator.New()}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(address)
}

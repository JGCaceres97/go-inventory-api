package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jgcaceres97/go-inventory-api/src/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.POST},
		AllowHeaders:     []string{echo.HeaderContentType},
		AllowOrigins:     []string{"http://127.0.0.1"},
	}))

	return e.Start(address)
}

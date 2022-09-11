package api

import (
	"log"
	"net/http"

	"github.com/jgcaceres97/go-inventory-api/src/encryption"
	"github.com/jgcaceres97/go-inventory-api/src/internal/api/dtos"
	"github.com/jgcaceres97/go-inventory-api/src/internal/models"
	"github.com/jgcaceres97/go-inventory-api/src/internal/service"
	"github.com/labstack/echo/v4"
)

type response struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, response{Message: "user already exists"})
		}

		return c.JSON(http.StatusInternalServerError, response{Message: "internal server error"})
	}

	return c.JSON(http.StatusCreated, nil)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response{Message: err.Error()})
	}

	user, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response{Message: "internal server error"})
	}

	token, err := encryption.SignedLoginToken(user)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response{Message: "internal server error"})
	}

	cookie := &http.Cookie{
		HttpOnly: true,
		Name:     "Authorization",
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Value:    token,
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func (a *API) AddProduct(c echo.Context) error {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response{Message: "unauthorized"})
	}

	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, response{Message: "unauthorized"})
	}

	email := claims["email"].(string)

	ctx := c.Request().Context()
	params := dtos.AddProduct{}

	err = c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response{Message: err.Error()})
	}

	product := models.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}

	err = a.serv.AddProduct(ctx, product, email)
	if err != nil {
		log.Println(err)

		if err == service.ErrInvalidPermission {
			return c.JSON(http.StatusForbidden, response{Message: "invalid permissions"})
		}

		return c.JSON(http.StatusInternalServerError, response{Message: "internal server error"})
	}

	return c.JSON(http.StatusCreated, nil)
}

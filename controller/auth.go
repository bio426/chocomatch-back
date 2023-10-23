package controller

import (
	"net/http"

	"github.com/bio426/chocomatch-back/service"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	service *service.Auth
}

func NewAuth(s *service.Auth) *Auth {
	ctr := &Auth{
		service: s,
	}

	return ctr
}

func (ctr Auth) Login(c echo.Context) error {
	ctr.service.Login()
	input := struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusNoContent)

}

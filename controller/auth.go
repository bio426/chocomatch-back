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

	token, err := ctr.service.Login(c.Request().Context(), input.Email, input.Password)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:  "authToken",
		Value: token,
	}
	c.SetCookie(cookie)

	return c.NoContent(http.StatusNoContent)

}

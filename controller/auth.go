package controller

import (
	"net/http"
	"time"

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
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	cookie := &http.Cookie{
		Name:     "authToken",
		Value:    token,
		Expires:  time.Now().Add(time.Second * 10),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)

	return c.NoContent(http.StatusNoContent)

}

func (ctr Auth) Register(c echo.Context) error {
	input := struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Phone    string `json:"phone"`
		Password string `json:"password" validate:"required"`
	}{}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctr.service.Register(c.Request().Context(), service.AuthRegisterArgs(input)); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)

}

func (ctr Auth) VerifyCookie(c echo.Context) error {
	cookie, err := c.Cookie("authToken")
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	output := struct {
		Value string `json:"value"`
	}{
		Value: cookie.Value,
	}

	return c.JSON(http.StatusOK, output)
}

func (ctr Auth) GetCookie(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "authToken",
		Value:    "xd2",
		Expires:  time.Now().Add(time.Second * 10),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)

	return c.NoContent(http.StatusNoContent)
}

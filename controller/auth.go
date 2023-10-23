package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var Auth = struct {
	Login func(c echo.Context) error
}{
	Login: func(c echo.Context) error {
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
	},
}

package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bio426/chocomatch-back/controller"
	"github.com/bio426/chocomatch-back/datasource"
)

func main() {
	config := datasource.InitConfig()
	pg, err := datasource.InitPostgres(config)
	if err != nil {
		log.Panic("",err, pg)
	}
	rds, err := datasource.InitRedis(config)
	if err != nil {
		log.Panic(err, rds)
	}
	e := echo.New()

	// Server config
	e.Debug = true
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}

	// Server middlewares
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Server routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api := e.Group("api")
	auth := api.Group("/auth")
	auth.POST("/login", controller.Auth.Login)

	// Run server
	e.Logger.Fatal(e.Start(":1323"))
}

// Validator setup
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

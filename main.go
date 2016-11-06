package main

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func init() {
	raven.SetDSN("https://xxxx@sentry.io/112319")
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.SetHTTPErrorHandler(func(err error, c echo.Context) {
		if _, ok := err.(*echo.HTTPError); !ok {
			raven.CaptureError(err, nil)
		}
		e.DefaultHTTPErrorHandler(err, c)
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/error", func(c echo.Context) error {
		return errors.New("something went wrong")
	})
	e.GET("/panic", func(c echo.Context) error {
		panic("panicked")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Run(standard.New(":1323"))

}

package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/shyfur-t", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Shyfur, Sumon")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
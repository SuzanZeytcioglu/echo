package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/echo", func(c echo.Context) error {
		data := c.FormValue("data")
		return c.String(http.StatusOK, data)
	})
	e.Start(":1323")
}

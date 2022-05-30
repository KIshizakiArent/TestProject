package echotest

import (
	"net/http"

	"github.com/labstack/echo"
)

func Echotest() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

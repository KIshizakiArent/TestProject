package echotest

import (
	"fmt"
	"net/http"
	"newProject/database"

	"github.com/labstack/echo"
)

func Echotest() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/database/", databaseHander)
	e.Logger.Fatal(e.Start(":1323"))
}

func databaseHander(c echo.Context) error {
	err := database.Connect()
	defer database.Close()
	if err != nil {
		return c.String(http.StatusUnauthorized, "connect failed")
	}
	err = database.Ping()
	if err != nil {
		return c.String(http.StatusUnauthorized, "ping failed")
	}

	rows, err := database.TableList()
	if err != nil {
		return c.String(http.StatusUnauthorized, "tableList failed")
	}
	defer rows.Close()

	var tableCount int
	for rows.Next() {
		err = rows.Scan()
		if err != nil {
			return c.String(http.StatusUnauthorized, "scan failed")
		}

		tableCount++
	}
	fmt.Println(tableCount)
	return c.String(http.StatusOK, "connected")
}

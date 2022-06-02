package echotest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type food struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

var foods []food

func init() {
	foods = make([]food, 0, 0)
}

func (f food) String() string {
	return fmt.Sprintf("%sの%s", f.Type, f.Name)
}

func getFood(c echo.Context) error {
	var foodList string
	for i := range foods {
		if i == 0 {
			foodList = fmt.Sprint(foods[i])
			continue
		}
		foodList = fmt.Sprint(foodList, ",", foods[i])
	}
	return c.String(http.StatusOK, foodList)
}

func addFood(c echo.Context) error {
	food := food{}

	err := c.Bind(&food)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	foods = append(foods, food)
	return c.String(http.StatusOK, fmt.Sprintf("%sを保管しました！", food))
}

func Echotest() {

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${host}${path} ${latency_human}" + "\n",
	}))

	e.Use(middleware.BasicAuth(func(userName, password string, c echo.Context) (bool, error) {
		if userName == "username" && password == "password" {
			fmt.Println("ログイン成功。")
			return true, nil
		}
		fmt.Println("ログイン失敗。")
		return false, nil
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello, World!"))
	})

	repositoryGroup := e.Group("/storage")

	repositoryGroup.POST("/food", addFood)
	repositoryGroup.GET("/food", getFood)

	e.Start(":8000")
}

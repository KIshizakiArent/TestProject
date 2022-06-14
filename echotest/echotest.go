package echotest

import (
	"net/http"
	oapi "newProject/api/dist"

	"github.com/labstack/echo/v4"
)

func Echotest() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 自動生成されたハンドラ登録関数にServerInterfaceを満たすserverを渡す
	handler := pingHandler{}
	oapi.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":1323"))
}

// 自動生成されたServerInterfaceを満たす構造体の素をつくる
type pingHandler struct{}

// ServerInterfaceのメソッドを実装していく
func (h pingHandler) GetPing(ctx echo.Context) error {
	// return値の第2引数に自動生成されたresponse型を使う
	return ctx.JSON(http.StatusOK, &oapi.Ping{
		Ping: "ping test ok",
	})
}

package echotest

import (
	"fmt"
	"net/http"
	oapi "newProject/api/dist"

	"github.com/go-playground/validator/v10"

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

	pingStruct := GetRequestTest()
	validate := validator.New()           //インスタンス生成
	errors := validate.Struct(pingStruct) //バリデーションを実行し、NGの場合、ここでエラーが返る。// return値の第2引数に自動生成されたresponse型を使う

	if errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Println("OK")
	}

	return ctx.JSON(http.StatusOK, oapi.Ping{
		Ping: pingStruct.Ping,
	})
}

type ping struct {
	Ping string `validate:"required"`
}

// 疑似RequestBodyからのGet結果をvalidate用structに詰め替え
func GetRequestTest() ping {
	return ping{Ping: "ping test"}
}

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Handler(ルーティングで動かす)
func show(c echo.Context) error {
	user := new(User)

	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func display(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()      // Instance
	e.GET("/user", show) // Routes
	e.POST("/users", display)
	e.Logger.Fatal(e.Start(":8000")) // Start server(内容を出力)
}

// 下記のアクセスでレスポンスが帰ってくる
// http://localhost:8000/user?name=yagi_eng&email=hoge@gmail.com

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoインスタンス作成
	e := echo.New()

	// 全てのリクエストで実行するミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]string{"Hello": "World"})
	// })
	// e.Logger.Fatal(e.Start(":8000"))

	e.GET("/hello", handler.MainPage())
	e.Start(":1323")
}

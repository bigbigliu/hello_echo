package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		fmt.Println("有一次新的http请求: ")
		return c.String(http.StatusOK, "Hello, World! - v7")
	})
	e.Logger.Fatal(e.Start(":9999"))
}

package main

import (
	"net/http"
	"os"

	"goconfig/config"

	"gorouter/logs"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		env := os.Getenv("configor_env")
		logs.Debug.Println("env:", env, "conn", config.Config.DB.Conn)
		conn := config.Config.DB.Conn
		return c.String(http.StatusOK, conn)
	})
	config.InitConfig()

	e.Start(":1113")
}

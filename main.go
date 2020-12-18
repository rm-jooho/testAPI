package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	corsConfig := middleware.DefaultCORSConfig
	e.Use(middleware.CORSWithConfig(corsConfig))
	e.POST("/liveApi/index.php/api/live_hook", handler)
	e.GET("/", handleRoot)
	err := e.Start(":80")
	if err != nil {
		panic(err)
	}
}

func handler(c echo.Context) error {
	m := echo.Map{}
	err := c.Bind(&m)
	if err != nil {
		fmt.Printf("handler ERR=%v\n", err)
		return c.NoContent(200)
	}
	fmt.Printf("handler m=%v\n", m)
	time.Sleep(5 * time.Second)
	fmt.Printf("handleer after sleep")
	return c.NoContent(200)
}

func handleRoot(c echo.Context) error {
	m := echo.Map{}
	m["result"] = true
	return c.JSON(200, m)
}

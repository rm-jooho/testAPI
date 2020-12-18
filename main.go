package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	corsConfig := middleware.DefaultCORSConfig
	e.Use(middleware.CORSWithConfig(corsConfig))
	e.POST("/liveApi/index.php/api/live_hook", handler)

	err := e.Start(":8011")
	if err != nil {
		panic(err)
	}
}

func handler(c echo.Context) error {
	m := echo.Map{}
	err := c.Bind(m)
	if err != nil {
		fmt.Printf("handler ERR=%v\n", err)
		return c.NoContent(200)
	}
	fmt.Printf("handleer m=%v\n", m)
	return c.NoContent(200)
}

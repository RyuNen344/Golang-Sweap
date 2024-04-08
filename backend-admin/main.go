package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("ADMIN_PORT")
	if "" == port {
		port = "8080"
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "backend-admin")
	})
	e.Logger.Fatal(e.Start(":" + port))
}

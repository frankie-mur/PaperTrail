package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", healthHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "up")
}

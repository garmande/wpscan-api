package main

import (
	"github.com/garmande/wpscan-api/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api/v1")
	api.POST("/scans", handler.CreateScan)
	api.GET("/scans", handler.ListScans)
	api.GET("/scans/:id", handler.GetScan)

	e.Logger.Fatal(e.Start(":8080"))
}

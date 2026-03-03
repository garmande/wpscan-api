package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const resultsDir = "/app/results"

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api/v1")
	api.GET("/results", getResults)

	e.Logger.Fatal(e.Start(":8080"))
}

func getResults(c echo.Context) error {
	data, err := os.ReadFile(resultsDir + "/latest.json")
	if err != nil {
		if os.IsNotExist(err) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "no scan results yet"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to read results"})
	}

	return c.JSONBlob(http.StatusOK, data)
}

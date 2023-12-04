package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		responseStr := ""
		for name, values := range c.Request().Header {
			// Loop over all values for the name.
			for _, value := range values {
				responseStr += fmt.Sprintf("%s: %s\n", name, value)
			}
		}
		return c.String(http.StatusOK, responseStr)
	})
	e.Logger.Fatal(e.Start(":1234"))
}

package main

import (
	"net/http"

	"github.com/alifudin-a/go-api-echo-psql/action"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome")
	})

	e.GET("/employee", action.GetEmployees)
	e.GET("/employee/:id", action.GetEmployee)
	e.POST("/employee", action.CreateEmployee)
	e.DELETE("/employee/:id", action.DeleteEmployee)
	e.PUT("/employee", action.UpdateEmployee)

	e.Logger.Fatal(e.Start(":8080"))
}

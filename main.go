package main

import (
	"github.com/alifudin-a/go-api-echo-psql/action"
	"github.com/alifudin-a/go-api-echo-psql/util"
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

	// Login route
	e.POST("/login", util.Login)

	// Unauthecnticatde route
	e.GET("/", util.Accessible)

	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &util.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", util.Restricted)

	// endpoint route for actions
	v1 := e.Group("/v1") // grouping
	v1.GET("/employee", action.GetEmployees)
	v1.GET("/employee/:id", action.GetEmployee)
	v1.POST("/employee", action.CreateEmployee)
	v1.DELETE("/employee/:id", action.DeleteEmployee)
	v1.PUT("/employee", action.UpdateEmployee)

	e.Logger.Fatal(e.Start(":8080"))
}

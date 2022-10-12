package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ladrillin/restaurantGo/db"
	"github.com/ladrillin/restaurantGo/routes"
)

func main() {

	server := echo.New()

	dbConn := db.DbConnection()
	defer dbConn.Close()

	routes.CustomerRoute(server)

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPatch,
		},
	}))

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.Logger.Fatal(server.Start(":8000"))
}

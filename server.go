package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	"github.com/ladrillin/restaurantGo/db"
	"github.com/ladrillin/restaurantGo/model"
)

func main() {

	dbConn := db.DbConnection()

	server := echo.New()

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

	server.POST("/api/order/add", func(c echo.Context) error {
		o := new(model.Order)
		if err := c.Bind(o); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO orders VALUES ($1, $2, $3)"
		res, err := dbConn.Query(sqlStatement, o.Id, o.CustomerId, o.Content)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, o)
		}
		return c.String(http.StatusOK, "ok")
	})

	server.Logger.Fatal(server.Start(":8000"))
}

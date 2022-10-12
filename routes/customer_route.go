package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ladrillin/restaurantGo/controllers"
)

func CustomerRoute(e *echo.Echo) {

	e.POST("/customers", controllers.CreateCustomer)
	e.GET("/customers/:id", controllers.GetCustomer)
	e.GET("/customers", controllers.GetAllCustomers)
	e.PUT("/customers/:id", controllers.UpdateCustomer)
	e.DELETE("/customers/:id", controllers.DeleteCustomer)
}

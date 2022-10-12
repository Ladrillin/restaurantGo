package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/ladrillin/restaurantGo/db"
	"github.com/ladrillin/restaurantGo/model"
	"github.com/ladrillin/restaurantGo/responses"
)

var validate = validator.New()
var dbConnection *sql.DB = db.DbConnection()

func CreateCustomer(c echo.Context) error {

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var customer model.Customer
	defer cancel()

	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, responses.CustomerResponse{
			Status:  http.StatusBadRequest,
			Message: "Binding Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if validationErr := validate.Struct(&customer); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.CustomerResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation Error",
			Data:    &echo.Map{"data": validationErr.Error()},
		})
	}

	newCustomer := model.Customer{
		Name:  customer.Name,
		Phone: customer.Phone,
		Email: customer.Email,
	}

	sqlStatement := `INSERT INTO customers
					(id, name, phone, email)
					VALUES (nextval('customers_sequence'), $1, $2, $3)`
	_, err := dbConnection.Exec(sqlStatement, newCustomer.Name, newCustomer.Phone, newCustomer.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.CustomerResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error while creating user",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusCreated, responses.CustomerResponse{
		Status:  http.StatusCreated,
		Message: "User created",
		Data:    &echo.Map{"data": newCustomer},
	})
}

func GetCustomer(c echo.Context) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	customerId := c.Param("id")
	var customer model.Customer
	defer cancel()

	sqlStatement := "SELECT * from customers WHERE id=$1"
	res := dbConnection.QueryRow(sqlStatement, customerId)

	err := res.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.CustomerResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error while GET user",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.CustomerResponse{
		Status:  http.StatusOK,
		Message: "GET responded successfully",
		Data:    &echo.Map{"data": customer},
	})

}

func UpdateCustomer(c echo.Context) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	customerId := c.Param("id")
	var customer model.Customer
	defer cancel()

	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, responses.CustomerResponse{
			Status:  http.StatusBadRequest,
			Message: "Binding Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	sqlStatement := `UPDATE customers SET name=$2, phone=$3, email=$4
					WHERE id=$1`
	_, err := dbConnection.Exec(sqlStatement, customerId, customer.Name, customer.Phone, customer.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.CustomerResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error while updating customer",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.CustomerResponse{
		Status:  http.StatusOK,
		Message: "PUT responded successfully",
		Data:    &echo.Map{"data": customer},
	})
}

func DeleteCustomer(c echo.Context) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	customerId := c.Param("id")
	defer cancel()

	sqlStatement := "DELETE FROM customers WHERE id=$1"
	_, err := dbConnection.Exec(sqlStatement, customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.CustomerResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error while deleting customer",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.CustomerResponse{
		Status:  http.StatusOK,
		Message: "User was deleted successfully",
		Data:    &echo.Map{"data": customerId},
	})
}

func GetAllCustomers(c echo.Context) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var customers []model.Customer

	sqlStatement := "SELECT * FROM customers"
	res, err := dbConnection.Query(sqlStatement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.CustomerResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error while get all users from DB",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	for res.Next() {
		var singleCustomer model.Customer
		err := res.Scan(&singleCustomer.Id, &singleCustomer.Name, &singleCustomer.Phone, &singleCustomer.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.CustomerResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error while getting all customer",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
		customers = append(customers, singleCustomer)
	}

	return c.JSON(http.StatusOK, responses.CustomerResponse{
		Status:  http.StatusOK,
		Message: "All users",
		Data:    &echo.Map{"data": customers},
	})
}

package app

import (
	"context"
	"example/employee-app/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *App) CreateEmployee(c *gin.Context) {
	employee := &models.Employee{}
	if err := c.BindJSON(employee); err != nil {
		fmt.Println(err.Error())
		return
	}

	createdEmployee, err := app.empRepo.CreateEmployee(context.Background(), employee)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, createdEmployee)
}

func (app *App) GetAllEmployees(c *gin.Context) {
	employees, err := app.empRepo.GetAllEmployees(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, employees)
}

func (app *App) GetEmployeeByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	employee, err := app.empRepo.GetEmployeeById(context.Background(), id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, employee)
}

func (app *App) DeleteEmployeeById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	err := app.empRepo.DeleteEmployee(context.Background(), id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{})
}

func (app *App) UpdateEmployee(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var newEmployee models.Employee

	if err := c.BindJSON(&newEmployee); err != nil {
		fmt.Println(err.Error())
		return
	}

	updated, err := app.empRepo.UpdateEmployee(context.Background(), id, &newEmployee)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, updated)
}

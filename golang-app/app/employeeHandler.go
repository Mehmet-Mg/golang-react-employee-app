package app

import (
	"example/employee-app/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *App) AddOneEmployee(c *gin.Context) {
	employee := &models.Employee{}
	if err := c.BindJSON(employee); err != nil {
		fmt.Println(err.Error())
		return
	}

	createdEmployee, err := app.repo.Create(*employee)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, createdEmployee)
}

func (app *App) GetAllEmployees(c *gin.Context) {
	employees, err := app.repo.All()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, employees)
}

func (app *App) GetEmployeeByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	employee, err := app.repo.GetById(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, employee)
}

func (app *App) DeleteEmployeeById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	err := app.repo.Delete(id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{})
}

func (app *App) UpdateEmployee(c *gin.Context) {
	var newEmployee models.Employee

	if err := c.BindJSON(&newEmployee); err != nil {
		fmt.Println(err.Error())
		return
	}

	updated, err := app.repo.Update(newEmployee.ID, newEmployee)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, updated)
}

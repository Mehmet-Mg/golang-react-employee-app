package app

import (
	"github.com/gin-gonic/gin"
)

func (app *App) Routes(router *gin.Engine) {
	router.GET("/employees", app.GetAllEmployees)
	router.GET("/employees/:id", app.GetEmployeeByID)
	router.POST("/employees", app.AddOneEmployee)
	router.PUT("/employees/:id", app.UpdateEmployee)
	router.DELETE("/employees/:id", app.DeleteEmployeeById)
}

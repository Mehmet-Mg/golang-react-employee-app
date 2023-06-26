package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// employee represents data about a record employee.
type employee struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Salary    float64 `json:"salary"`
}

// employees slice to seed record album data.
var employees = []employee{
	{ID: "1", FirstName: "John", LastName: "Coltrane", Salary: 56.99},
	{ID: "2", FirstName: "Gerry", LastName: "Mulligan", Salary: 17.99},
	{ID: "3", FirstName: "Sarah", LastName: "Vaughan", Salary: 39.99},
}

var lastId = 4

func main() {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", getEmployeeByID)
	router.POST("/employees", postEmployee)
	router.PUT("/employees/:id", updateEmployee)
	router.DELETE("/employees/:id", deleteEmployeeById)

	router.Run("localhost:8080")
}

// getEmployees responds with the list of all albums as JSON.
func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

// postEmployee adds an album from JSON received in the request body.
func postEmployee(c *gin.Context) {
	var newEmployee employee

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newEmployee); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Add the new album to the slice.
	employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

// getEmployeeByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range employees {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
}

func deleteEmployeeById(c *gin.Context) {
	id := c.Param("id")

	newEmployeeList := []employee{}

	for _, k := range employees {
		if k.ID != id {
			newEmployeeList = append(newEmployeeList, k)
		}
	}

	employees = newEmployeeList
}

// postEmployee adds an album from JSON received in the request body.
func updateEmployee(c *gin.Context) {
	var newEmployee employee
	id := c.Param("id")

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newEmployee); err != nil {
		fmt.Println(err.Error())
		return
	}

	// update
	for i, k := range employees {
		if k.ID == id {
			employees[i] = newEmployee
		}
	}

	c.IndentedJSON(http.StatusCreated, newEmployee)
}

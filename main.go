package main

import (
	"errors"
	"fmt"
	"main/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func handleGetEmployees(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.IndentedJSON(http.StatusOK, employees)
}

func handleOneEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error // SELECT * FROM Employee WHERE id=id
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func handleNewEmployees(c *gin.Context) {

	var employee data.Employee
	// err:= c.BindJSON(&employee)
	// if(err != nil)
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id = 0
	err := data.DB.Create(&employee).Error
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, employee)
	}
}

func handleStartPage(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body><h1>Hej hopp</h1><ul><li>Test</li><li>Bla</li></ul></body></html>"))
}

var config Config

func main() {
	readConfig(&config)
	fmt.Println(config)
	//	emps := data.GetAl
	data.InitDatabase(config.Database.File)

	fmt.Println(config)
	//	emps := data.GetAllEmployees()

	// GET
	// POST
	// PUT
	r := gin.Default()

	r.GET("/", handleStartPage)
	r.GET("/api/employee", handleGetEmployees)
	r.GET("/api/employee/:id", handleOneEmployee)
	r.POST("/api/employee", handleNewEmployees)

	r.Run(":8080")

	//fmt.Println("312123213")
}

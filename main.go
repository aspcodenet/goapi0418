package main

import (
	"errors"
	"fmt"
	"main/data"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// hello

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

func apiEmployeeUpdateById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		if err := c.BindJSON(&employee); err != nil {
			return
		}
		employee.Id, _ = strconv.Atoi(id)
		data.DB.Save(&employee)
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func apiEmployeeDeleteById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		data.DB.Delete(&employee)
		c.IndentedJSON(http.StatusNoContent, employee)
	}
}

func handleStartPage(c *gin.Context) {
	name, _ := os.Hostname()
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body><h1>Hej hopp2</h1><ul><li>Dator:"+name+"</li><li>Bla</li></ul></body></html>"))
}

// hej

var config Config

func main() {
	readConfig(&config)
	fmt.Println(config)
	//	emps := data.GetAl
	data.InitDatabase(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

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
	r.PUT("/api/employee/:id", apiEmployeeUpdateById)
	r.PATCH("/api/employee/:id", apiEmployeeUpdateById)

	r.DELETE("/api/employee/:id", apiEmployeeDeleteById)

	r.Run(":8080")

	//fmt.Println("312123213")
}

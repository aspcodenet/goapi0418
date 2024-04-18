package main

import (
	"fmt"
	"main/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGetEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func handleOneEmployee(c *gin.Context) {
	id := c.Param("id")

	numId, _ := strconv.Atoi(id)

	employee := data.GetEmployee(numId)
	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func handleNewEmployees(c *gin.Context) {
	// TODO Add new
	// försöka få fram den JSON Employee som man skickat in
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	data.SaveEmployee(employee)
	c.IndentedJSON(http.StatusCreated, employee)
}

func handleStartPage(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body><h1>Hej hopp</h1><ul><li>Test</li><li>Bla</li></ul></body></html>"))
}

func main() {
	data.Init()
	//	emps := data.GetAllEmployees()

	// GET
	// POST
	// PUT
	r := gin.Default()

	r.GET("/", handleStartPage)
	r.GET("/api/employee", handleGetEmployees)
	r.GET("/api/employee/:id", handleOneEmployee)
	r.POST("/api/employee", handleNewEmployees)

	r.Run(":8081")

	fmt.Println("312123213")
}

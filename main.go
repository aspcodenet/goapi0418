package main

import (
	"main/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func handleNewEmployees(c *gin.Context) {
	// TODO Add new
	// försöka få fram den JSON Employee som man skickat in
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	data.SaveEmployee(employee)
	c.JSON(http.StatusCreated, employee)
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
	r.POST("/api/employee", handleNewEmployees)

	r.Run(":8081")

	// for _, emp := range emps {
	// 	fmt.Println(emp)
	// }
}

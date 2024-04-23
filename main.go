package main

import (
	"main/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetEmployees(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.IndentedJSON(http.StatusOK, employees)
}

func handleOneEmployee(c *gin.Context) {
	// id := c.Param("id")

	// numId, _ := strconv.Atoi(id)

	// employee := data.GetEmployee(numId)
	// if employee == nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	// } else {
	// 	c.JSON(http.StatusOK, employee)
	// }
}

func handleNewEmployees(c *gin.Context) {
	// // TODO Add new
	// // försöka få fram den JSON Employee som man skickat in
	// var employee data.Employee
	// if err := c.BindJSON(&employee); err != nil {
	// 	return
	// }
	// data.SaveEmployee(employee)
	// c.IndentedJSON(http.StatusCreated, employee)
}

func handleStartPage(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body><h1>Hej hopp</h1><ul><li>Test</li><li>Bla</li></ul></body></html>"))
}

func main() {
	data.InitDatabase()
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

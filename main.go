package main

import (
	"fmt"
	"main/data"
)

func main() {
	data.Init()
	emps := data.GetAllEmployees()

	for _, emp := range emps {
		fmt.Println(emp)
	}
}

package data

var employees []Employee

func GetAllEmployees() []Employee {
	return employees
}

func Init() {
	employees = append(employees, Employee{Age: 50, Namn: "Stefan", City: "Test"})
	employees = append(employees, Employee{Age: 14, Namn: "Oliver", City: "Test"})
	employees = append(employees, Employee{Age: 20, Namn: "Josefine", City: "Test"})
}

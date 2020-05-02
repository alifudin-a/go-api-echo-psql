package model

// Employee : struct model for employee
type Employee struct {
	ID     string `json:"id" validate:"required"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

// Employees : struct for list all employee
type Employees struct {
	Employees []Employee `json:"employees"`
}

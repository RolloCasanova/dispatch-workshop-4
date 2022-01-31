package model

// Employees type is an alias for a slice of Employee.
type Employees []Employee

// Employee struct represents a single employee.
type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

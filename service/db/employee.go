package db

import (
	"log"

	errz "github.com/RolloCasanova/dispatch-workshop-4/errors"
	"github.com/RolloCasanova/dispatch-workshop-4/model"
)

// EmployeeMap is an alias for a map of employees.
type EmployeeMap map[int]model.Employee

// employeesOrder is an auxiliary function that helps to sort employeeData by their ID
// as we are using a map we can't ensure the order of the keys is preserved.
var employeesOrder []int = []int{1, 2}

// db is a sample data to be used in the service as a placeholder for the real data.
var db EmployeeMap = map[int]model.Employee{
	1: {
		ID:      1,
		Name:    "1",
		Email:   "1",
		Phone:   "1",
		Address: "1",
	},
	2: {
		ID:      2,
		Name:    "2",
		Email:   "2",
		Phone:   "2",
		Address: "2",
	},
}

// employeeDBService struct consists of the data to be used in the service.
type employeeDBService struct {
	data EmployeeMap
}

// New returns a new EmployeeService instance.
func New(em EmployeeMap) employeeDBService {
	if em == nil {
		em = db
	}

	return employeeDBService{
		data: em,
	}
}

// GetAllEmployees returns all employees data.
// Will retrieve data from the db and update the cache.
func (es employeeDBService) GetAllEmployees() (model.Employees, error) {
	log.Println("In db service - GetAllEmployees")

	if err := es.dataValidation(); err != nil {
		return nil, err
	}

	// convert data from map to an slice of Employees
	employees := make(model.Employees, 0, len(es.data))

	// preserve the order
	for _, id := range employeesOrder {
		employees = append(employees, es.data[id])
	}

	return employees, nil
}

// GetEmployeeByID returns an employee by its ID.
// first look at redis, and if it doesn't exist, look at the db.
func (es employeeDBService) GetEmployeeByID(id int) (*model.Employee, error) {
	log.Println("In db service - GetEmployeeByID")

	if err := es.dataValidation(); err != nil {
		return nil, err
	}

	// find the employee in the data
	employee, ok := es.data[id]
	if !ok {
		return nil, errz.ErrNotFound
	}

	return &employee, nil
}

// CreateEmployee creates a new employee and add it to the db
// if it already exists it won't be overwritten.
// it also adds the employee to the cache.
func (es employeeDBService) CreateEmployee(e model.Employee) error {
	log.Println("In db service - CreateEmployee")

	if err := es.dataValidation(); err != nil {
		return err
	}

	// special handling if employee already exists
	if _, ok := es.data[e.ID]; ok {
		return errz.ErrEmployeeAlreadyExists
	}

	// add employee
	es.data[e.ID] = e

	// update employees order
	employeesOrder = append(employeesOrder, e.ID)

	return nil
}

// dataValidation is an auxiliary function that checks if the data has beem initialized or if it is empty
// returns a matching ServiceError if any of these conditions are met.
func (es *employeeDBService) dataValidation() error {
	log.Println("In db service - dataValidation")

	// special handling if data is nil
	if es.data == nil {
		return errz.ErrDataNotInitialized
	}

	// special handling if data is empty
	if len(es.data) == 0 {
		return errz.ErrEmptyData
	}

	return nil
}

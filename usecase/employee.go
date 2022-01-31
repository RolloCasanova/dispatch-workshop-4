package usecase

import (
	"fmt"
	"log"

	"github.com/RolloCasanova/dispatch-workshop-4/model"
	"github.com/go-redis/redis"
)

// EmployeeDBService is the interface that wraps the DB service's methods
// GetAllEmployees, GetEmployeeByID, and CreateEmployee.
type EmployeeDBService interface {
	GetAllEmployees() (model.Employees, error)
	GetEmployeeByID(id int) (*model.Employee, error)
	CreateEmployee(e model.Employee) error
}

// employeeRedisService is the interface that wraps the Redis service's methods
// UpsertEmployee, ReadEmployee and DeleteEmployeeByID
type employeeRedisService interface {
	UpsertEmployee(e model.Employee) error
	UpsertEmployees(e model.Employees) error
	ReadEmployee(id int) (*model.Employee, error)
	DeleteEmployeeByID(id int) error
}

// employeeUsecase implements EmployeeService interface.
type employeeUsecase struct {
	db EmployeeDBService
	rd employeeRedisService
}

// New returns a new EmployeeUsecase instance.
func New(db EmployeeDBService, rd employeeRedisService) employeeUsecase {
	log.Println("In usecase - NewEmployeeUsecase")

	return employeeUsecase{
		db: db,
		rd: rd,
	}
}

// GetAllEmployees get all employees from the db service and update the cache.
func (eu employeeUsecase) GetAllEmployees() (model.Employees, error) {
	log.Println("In usecase - GetAllEmployees")

	employees, err := eu.db.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	if err := eu.rd.UpsertEmployees(employees); err != nil {
		return nil, err
	}

	return employees, nil
}

// GetEmployeeByID calls the service to returns an employee by its ID.
// If the employee is not found in the cache, it will be read from the db service.
func (eu employeeUsecase) GetEmployeeByID(id int) (*model.Employee, error) {
	log.Println("In usecase - GetEmployeeByID")

	// get the employee from the cache
	employee, err := eu.rd.ReadEmployee(id)
	if err != nil {
		if err == redis.Nil {
			log.Println("employee not found in cache")
		} else {
			return nil, fmt.Errorf("reading employee from cache: %v", err)
		}
	}

	// if the employee is not found in the cache, get it from the db service
	if employee == nil {
		employee, err = eu.db.GetEmployeeByID(id)
		if err != nil {
			return nil, fmt.Errorf("getting employee from db service: %v", err)
		}

		// update the cache
		if eu.rd.UpsertEmployee(*employee) != nil {
			return nil, fmt.Errorf("writing employee to cache: %v", err)
		}
	}

	return employee, nil
}

// CreateEmployee calls the db service to create an employeze.
// if success, then it will try to write the employee to the cache.
func (eu employeeUsecase) CreateEmployee(e model.Employee) error {
	log.Println("In usecase - CreateEmployee")

	if err := eu.db.CreateEmployee(e); err != nil {
		return fmt.Errorf("creating employee in db: %v", err)
	}

	if err := eu.rd.UpsertEmployee(e); err != nil {
		return fmt.Errorf("writing employee to cache: %v", err)
	}

	return nil
}

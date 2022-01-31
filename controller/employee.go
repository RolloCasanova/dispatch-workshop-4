package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	errz "github.com/RolloCasanova/dispatch-workshop-4/errors"
	"github.com/RolloCasanova/dispatch-workshop-4/model"
	"github.com/gorilla/mux"
)

// EmployeeUsecase is the interface that wraps the EmployeeUsecase's methods
// GetEmployeeByID, GetAllEmployees, and CreateEmployee.
type EmployeeUsecase interface {
	GetEmployeeByID(id int) (*model.Employee, error)
	GetAllEmployees() (model.Employees, error)
	CreateEmployee(e model.Employee) error
}

// employeeController implements EmployeeUsecase interface.
type employeeController struct {
	usecase EmployeeUsecase
}

// New returns a new EmployeeController instance.
func New(uc EmployeeUsecase) employeeController {
	return employeeController{
		usecase: uc,
	}
}

// GetEmployeeByID returns an employee by its ID.
func (ec employeeController) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller - GetEmployeeByID")

	// extract the path parameters
	vars := mux.Vars(r)

	// convert the id param into an int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid id: %v", err)

		log.Printf("converting id param into an int: %v", err)
		return
	}

	// get the employee from the usecase
	employee, err := ec.usecase.GetEmployeeByID(id)
	if err != nil {
		switch {
		case errors.Is(err, errz.ErrNotFound), errors.Is(err, errz.ErrEmptyData):
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "employee not found")

		case errors.Is(err, errz.ErrDataNotInitialized):
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "data not initialized")

			log.Printf("getting employee: %v", err)
		}
	}

	if employee == nil {
		log.Println("no employee found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "no employee found")

		return
	}

	jsonData, err := json.Marshal(employee)
	if err != nil {
		log.Println("error marshalling employees")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, "error marshalling employees")
	}

	// this is fine
	log.Printf("employee found: %+v", employee)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// GetAllEmployees calls the usecase to return all employees.
func (ec employeeController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller - GetAllEmployees")

	// get all employees from the usecase
	employees, err := ec.usecase.GetAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting employees")

		log.Printf("getting all employees: %v", err)
	}

	// special handling if employees is empty
	if len(employees) == 0 {
		log.Println("no employees found")
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintln(w, "no employees found")
		return
	}

	jsonData, err := json.Marshal(employees)
	if err != nil {
		log.Println("error marshalling employees")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, "error marshalling employees: %v\n", err)
	}

	// this is fine
	log.Printf("employees found: %+v\n", employees)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// CreateEmployee calls the usecase to create a new employee.
// It receives a JSON payload with the employee data.
func (ec employeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller - CreateEmployee")

	// decode the request body into a new employee
	employee := model.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	// not recomented to handle this way
	if err != nil {
		log.Printf("r.Body: %+v\n", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request body")

		log.Printf("decoding request body: %v", err)
	}

	// create the employee in the usecase
	err = ec.usecase.CreateEmployee(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating employee: %v", err)
		log.Printf("creating employee: %v", err)

		return
	}

	// this is fine
	log.Printf("employee created: %+v\n", employee)
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "employee created")
}

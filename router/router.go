package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// EmployeeController is the interface that wraps the controller's methods
// GetEmployeeByID, GetAllEmployees, and CreateEmployee.
type EmployeeController interface {
	GetEmployeeByID(w http.ResponseWriter, r *http.Request)
	GetAllEmployees(w http.ResponseWriter, r *http.Request)
	CreateEmployee(w http.ResponseWriter, r *http.Request)
}

// Setup returns a router instance
func Setup(c EmployeeController) *mux.Router {
	r := mux.NewRouter()

	// versioning api
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// the endpoints
	v1.HandleFunc("/employees", c.GetAllEmployees).
		Methods(http.MethodGet).Name("GetAllEmployees")

	v1.HandleFunc("/employees/{id}", c.GetEmployeeByID).
		Methods(http.MethodGet).Name("GetEmployeeByID")

	v1.HandleFunc("/employees", c.CreateEmployee).
		Methods(http.MethodPost).Name("CreateEmployee")

	return r
}

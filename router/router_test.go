package router

import (
	"net/http"
	"testing"

	"github.com/RolloCasanova/dispatch-workshop-4/mocks"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name      string
		routeName string
		path      string
		methods   []string
	}{
		{
			name:      "GetAllEmployees route",
			routeName: "GetAllEmployees",
			path:      "/api/v1/employees",
			methods:   []string{http.MethodGet},
		},
		{
			name:      "GetEmployeeByID route",
			routeName: "GetEmployeeByID",
			path:      "/api/v1/employees/{id}",
			methods:   []string{http.MethodGet},
		},
		{
			name:      "CreateEmployee route",
			routeName: "CreateEmployee",
			path:      "/api/v1/employees",
			methods:   []string{http.MethodPost},
		},
		// {
		// 	name:      "Random route",
		// 	routeName: "Random",
		// 	path:      "/api/v1/random",
		// 	methods:   []string{http.MethodTrace},
		// },
	}

	r := Setup(&mocks.EmployeeController{})

	for _, tt := range tests {
		// get the registered route
		route := r.Get(tt.routeName)

		if route == nil {
			t.Errorf("%s route is not registered - should be registered", tt.routeName)
			t.FailNow()
		}

		name := route.GetName()
		if name != tt.routeName {
			t.Errorf("route name is: %s - you have: %s", name, tt.routeName)
			t.FailNow()
		}

		path, _ := route.GetPathTemplate()
		if path != tt.path {
			t.Errorf("route path is: %s - you have: %s", path, tt.path)
			t.FailNow()
		}

		methods, _ := route.GetMethods()
		// deep equal check
		assert.EqualValues(t, methods, tt.methods, "route methods are not equal")
	}
}

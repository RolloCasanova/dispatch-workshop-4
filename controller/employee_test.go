package controller

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/RolloCasanova/dispatch-workshop-4/errors"
	"github.com/RolloCasanova/dispatch-workshop-4/mocks"
	"github.com/RolloCasanova/dispatch-workshop-4/model"
	"github.com/gorilla/mux"
)

func TestNew(t *testing.T) {
	type args struct {
		uc EmployeeUsecase
	}
	tests := []struct {
		name string
		args args
		want employeeController
	}{
		{
			name: "Happy path - create a new controller",
			args: args{
				uc: &mocks.EmployeeUsecase{},
			},
			want: employeeController{
				usecase: &mocks.EmployeeUsecase{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.uc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	anyEmployee = &model.Employee{
		ID:   1,
		Name: "John Doe",
	}

	anyPathParams = map[string]string{
		"id": "1",
	}

	anyID = 1
)

func Test_employeeController_GetEmployeeByID(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		pathParams  map[string]string
		callUsecase bool
		ucEmployee  *model.Employee
		ucError     error
		want        int
	}{
		{
			name:        "Happy path - get an employee by its ID",
			id:          1,
			pathParams:  anyPathParams,
			callUsecase: true,
			ucEmployee:  anyEmployee,
			want:        http.StatusOK,
		},
		{
			name: "Should fail - no ID in the path",
			want: http.StatusBadRequest,
		},
		{
			name:        "Should fail - error not found while getting the employee",
			id:          anyID,
			pathParams:  anyPathParams,
			callUsecase: true,
			ucError:     errors.ErrNotFound,
			want:        http.StatusNotFound,
		},
		{
			name:        "Should fail - error data not initalized while getting the employee",
			id:          anyID,
			pathParams:  anyPathParams,
			callUsecase: true,
			ucError:     errors.ErrDataNotInitialized,
			want:        http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		uc := &mocks.EmployeeUsecase{}

		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/employees/:id", nil)
			req = mux.SetURLVars(req, tt.pathParams)

			if tt.callUsecase {
				uc.On("GetEmployeeByID", tt.id).Return(tt.ucEmployee, tt.ucError)
			}

			ec := employeeController{
				usecase: uc,
			}

			ec.GetEmployeeByID(rw, req)

			if rw.Code != tt.want {
				t.Errorf("employeeController.GetEmployeeByID() = %v, want %v", rw.Code, tt.want)
			}

			uc.AssertExpectations(t)
		})
	}
}

func Test_employeeController_GetAllEmployees(t *testing.T) {
	type fields struct {
		usecase EmployeeUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := employeeController{
				usecase: tt.fields.usecase,
			}
			ec.GetAllEmployees(tt.args.w, tt.args.r)
		})
	}
}

func Test_employeeController_CreateEmployee(t *testing.T) {
	type fields struct {
		usecase EmployeeUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := employeeController{
				usecase: tt.fields.usecase,
			}
			ec.CreateEmployee(tt.args.w, tt.args.r)
		})
	}
}

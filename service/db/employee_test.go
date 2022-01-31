package db

import (
	"reflect"
	"testing"

	"github.com/RolloCasanova/dispatch-workshop-4/errors"
	"github.com/RolloCasanova/dispatch-workshop-4/model"
)

var (
	anyEmployeeMap = EmployeeMap{
		1: {
			ID:   1,
			Name: "John Doe",
		},
		2: {
			ID:   2,
			Name: "Jane Doe",
		},
		3: {
			ID:   3,
			Name: "John Smith",
		},
	}

	anyEmployeeOrder = []int{1, 2, 3}

	anyEmployees = model.Employees{
		{
			ID:   1,
			Name: "John Doe",
		},
		{
			ID:   2,
			Name: "Jane Doe",
		},
		{
			ID:   3,
			Name: "John Smith",
		},
	}
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		em   EmployeeMap
		want employeeDBService
	}{
		{
			name: "Happy path - sending an initialized map",
			em:   anyEmployeeMap,
			want: employeeDBService{
				data: anyEmployeeMap,
			},
		},
		{
			name: "Should return default map if no map is sent",
			want: employeeDBService{
				data: db,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.em); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeDBService_GetAllEmployees(t *testing.T) {
	type fields struct {
		data EmployeeMap
	}
	tests := []struct {
		name          string
		fields        fields
		employeeOrder []int
		want          model.Employees
		dbError       error
	}{
		{
			name: "Happy path - should return all employees",
			fields: fields{
				data: anyEmployeeMap,
			},
			employeeOrder: anyEmployeeOrder,
			want:          anyEmployees,
		},
		{
			name:    "Should fail - data has not been initialized",
			dbError: errors.ErrDataNotInitialized,
		},
		{
			name:    "Should fail - data is empty",
			fields:  fields{data: EmployeeMap{}},
			dbError: errors.ErrEmptyData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employeesOrder = tt.employeeOrder

			es := employeeDBService{
				data: tt.fields.data,
			}
			got, err := es.GetAllEmployees()

			if err != tt.dbError {
				t.Errorf("employeeDBService.GetAllEmployees() error = %v, wantErr %v", err, tt.dbError)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeDBService.GetAllEmployees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeDBService_GetEmployeeByID(t *testing.T) {
	type fields struct {
		data EmployeeMap
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := employeeDBService{
				data: tt.fields.data,
			}
			got, err := es.GetEmployeeByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("employeeDBService.GetEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeDBService.GetEmployeeByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeDBService_CreateEmployee(t *testing.T) {
	type fields struct {
		data EmployeeMap
	}
	type args struct {
		e model.Employee
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := employeeDBService{
				data: tt.fields.data,
			}
			if err := es.CreateEmployee(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("employeeDBService.CreateEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_employeeDBService_dataValidation(t *testing.T) {
	type fields struct {
		data EmployeeMap
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := &employeeDBService{
				data: tt.fields.data,
			}
			if err := es.dataValidation(); (err != nil) != tt.wantErr {
				t.Errorf("employeeDBService.dataValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package usecase

import (
	"reflect"
	"testing"

	"github.com/RolloCasanova/dispatch-workshop-4/errors"
	"github.com/RolloCasanova/dispatch-workshop-4/mocks"
	"github.com/RolloCasanova/dispatch-workshop-4/model"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		db EmployeeDBService
		rd EmployeeRedisService
	}
	tests := []struct {
		name string
		args args
		want employeeUsecase
	}{
		{
			name: "New employeeUsecase",
			args: args{
				db: &mocks.EmployeeDBService{},
				rd: &mocks.EmployeeRedisService{},
			},
			want: employeeUsecase{
				db: &mocks.EmployeeDBService{},
				rd: &mocks.EmployeeRedisService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db, tt.args.rd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	anyEmployees = model.Employees{
		{
			ID:   1,
			Name: "John Doe",
		},
		{
			ID:   2,
			Name: "Jane Doe",
		},
	}
)

func Test_employeeUsecase_GetAllEmployees(t *testing.T) {
	tests := []struct {
		name               string
		dbEmployeeResponse model.Employees
		dbEmployeeError    error
		rdEmployeeError    error
		callRdService      bool
	}{
		{
			name:               "Happy path - Get all employees",
			dbEmployeeResponse: anyEmployees,
			callRdService:      true,
		},
		{
			name:            "Should fail - db service error",
			dbEmployeeError: errors.ErrNotFound,
		},
		{
			name:            "Should fail - rd service error",
			callRdService:   true,
			rdEmployeeError: errors.ErrDataNotInitialized,
		},
	}

	for _, tt := range tests {
		db := &mocks.EmployeeDBService{}
		rd := &mocks.EmployeeRedisService{}

		t.Run(tt.name, func(t *testing.T) {
			eu := employeeUsecase{db, rd}

			db.On("GetAllEmployees").Return(tt.dbEmployeeResponse, tt.dbEmployeeError)

			employees, err := eu.db.GetAllEmployees()

			if tt.callRdService {
				rd.On("UpsertEmployees", tt.dbEmployeeResponse).Return(tt.rdEmployeeError)
				err = eu.rd.UpsertEmployees(tt.dbEmployeeResponse)
			}

			got, err := eu.GetAllEmployees()

			if tt.dbEmployeeError != nil {
				assert.Equal(t, tt.dbEmployeeResponse, got)
				assert.EqualError(t, err, tt.dbEmployeeError.Error())
			}

			if tt.rdEmployeeError != nil {
				assert.EqualError(t, err, tt.rdEmployeeError.Error())
			}

			assert.Equal(t, employees, got)
			db.AssertExpectations(t)
			rd.AssertExpectations(t)
		})
	}
}

func Test_employeeUsecase_GetEmployeeByID(t *testing.T) {
	type fields struct {
		db EmployeeDBService
		rd EmployeeRedisService
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
			eu := employeeUsecase{
				db: tt.fields.db,
				rd: tt.fields.rd,
			}
			got, err := eu.GetEmployeeByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("employeeUsecase.GetEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeUsecase.GetEmployeeByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeUsecase_CreateEmployee(t *testing.T) {
	type fields struct {
		db EmployeeDBService
		rd EmployeeRedisService
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
			eu := employeeUsecase{
				db: tt.fields.db,
				rd: tt.fields.rd,
			}
			if err := eu.CreateEmployee(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("employeeUsecase.CreateEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

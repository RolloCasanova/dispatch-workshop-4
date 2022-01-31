package errors

import (
	errs "errors"
)

type ServiceError error

var (
	ErrNotFound              ServiceError = errs.New("employee not found")
	ErrEmptyData             ServiceError = errs.New("data is empty")
	ErrDataNotInitialized    ServiceError = errs.New("data not initialized")
	ErrEmployeeAlreadyExists ServiceError = errs.New("employee already exists")
)

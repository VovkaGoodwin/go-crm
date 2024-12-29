package app_errors

import "github.com/pkg/errors"

var ErrDepartmentWithThisNameExists = errors.New("department with this name already exists")

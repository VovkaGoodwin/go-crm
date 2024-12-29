package dto

import "crm-backend/internal/rybakcrm/app/domain/models"

type CreateDepartmentRequest struct {
	Name        string
	Description string
}

type CreateDepartmentResponse struct {
	Department *models.Department
}

type GetAllDepartmentsRequest struct{}

type GetAllDepartmentsResponse struct {
	Departments []*models.Department
}

type GetDepartmentByIdRequest struct {
	DepartmentId int
}
type GetDepartmentByIdResponse struct {
	Department *models.Department
}

type UpdateDepartmentRequest struct {
	DepartmentId int
	Name         string
	Description  string
}
type UpdateDepartmentResponse struct {
	Department *models.Department
}

type DeleteDepartmentRequest struct {
	DepartmentId int
}

type DeleteDepartmentResponse struct {
	Success bool
}

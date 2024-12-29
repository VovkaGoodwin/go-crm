package interactors

import (
	"context"
	"crm-backend/internal/rybakcrm/app/application/dto"
	"crm-backend/internal/rybakcrm/app/domain/models"
	"crm-backend/internal/rybakcrm/app/domain/service"
	"crm-backend/internal/rybakcrm/config"
	"log/slog"
)

type DepartmentInteractor struct {
	cfg     *config.Config
	log     *slog.Logger
	service *service.DepartmentService
}

func NewDepartmentInteractor(
	cfg *config.Config,
	log *slog.Logger,
	service *service.DepartmentService,
) *DepartmentInteractor {
	return &DepartmentInteractor{
		cfg:     cfg,
		log:     log,
		service: service,
	}
}

func (d *DepartmentInteractor) CreateDepartment(
	ctx context.Context,
	request *dto.CreateDepartmentRequest,
) (
	*dto.CreateDepartmentResponse,
	error,
) {

	department, err := d.service.CreateDepartment(ctx, request.Name, request.Description)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateDepartmentResponse{
		Department: department,
	}

	return response, nil
}

func (
	d *DepartmentInteractor,
) GetAllDepartments(
	ctx context.Context,
	req *dto.GetAllDepartmentsRequest,
) (*dto.GetAllDepartmentsResponse, error) {

	departments, err := d.service.GetAllDepartments(ctx)
	if err != nil {
		return nil, err
	}

	response := &dto.GetAllDepartmentsResponse{
		Departments: departments,
	}

	return response, nil
}

func (d *DepartmentInteractor) GetDepartmentById(ctx context.Context, req *dto.GetDepartmentByIdRequest) (*dto.GetDepartmentByIdResponse, error) {
	department, err := d.service.GetDepartment(ctx, req.DepartmentId)
	if err != nil {
		return nil, err
	}

	response := &dto.GetDepartmentByIdResponse{
		Department: department,
	}

	return response, nil
}

func (d *DepartmentInteractor) UpdateDepartment(ctx context.Context, request *dto.UpdateDepartmentRequest) (*dto.UpdateDepartmentResponse, error) {
	department := &models.Department{
		Id:          request.DepartmentId,
		Name:        request.Name,
		Description: request.Description,
	}

	department, err := d.service.UpdateDepartment(ctx, department)

	if err != nil {
		return nil, err
	}

	response := &dto.UpdateDepartmentResponse{
		Department: department,
	}

	return response, nil
}

func (d *DepartmentInteractor) DeleteDepartment(ctx context.Context, request *dto.DeleteDepartmentRequest) (*dto.DeleteDepartmentResponse, error) {
	success, err := d.service.DeleteDepartment(ctx, request.DepartmentId)

	response := &dto.DeleteDepartmentResponse{
		Success: success,
	}

	return response, err
}

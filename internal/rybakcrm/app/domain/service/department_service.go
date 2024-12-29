package service

import (
	"context"
	"crm-backend/pkg/app_errors"

	"crm-backend/internal/rybakcrm/app/domain/models"
	"crm-backend/internal/rybakcrm/app/domain/repository"
)

type DepartmentService struct {
	departmentRepository repository.DepartmentRepository
}

func NewDepartmentService(departmentRepository repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		departmentRepository: departmentRepository,
	}
}

func (d *DepartmentService) CreateDepartment(ctx context.Context, name, description string) (*models.Department, error) {

	dep := &models.Department{
		Name:        name,
		Description: description,
	}

	result, err := d.departmentRepository.CreateDepartment(ctx, dep)

	return result, err
}

func (d *DepartmentService) GetAllDepartments(
	ctx context.Context,
) ([]*models.Department, error) {
	return d.departmentRepository.GetAllDepartments(ctx)
}

func (d *DepartmentService) GetDepartment(ctx context.Context, id int) (*models.Department, error) {
	return d.departmentRepository.GetDepartment(ctx, id)
}

func (d *DepartmentService) UpdateDepartment(ctx context.Context, department *models.Department) (*models.Department, error) {
	depExists, err := d.departmentRepository.CheckDepartmentExists(ctx, department)
	if err != nil {
		return nil, err
	}

	if depExists {
		return nil, app_errors.ErrDepartmentWithThisNameExists
	}

	return d.departmentRepository.UpdateDepartment(ctx, department)
}

func (d *DepartmentService) DeleteDepartment(ctx context.Context, id int) (bool, error) {
	return d.departmentRepository.DeleteDepartment(ctx, id)
}

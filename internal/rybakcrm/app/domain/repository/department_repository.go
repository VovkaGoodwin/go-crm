package repository

import (
	"context"
	"crm-backend/internal/rybakcrm/app/domain/models"
)

type DepartmentRepository interface {
	CreateDepartment(ctx context.Context, dto *models.Department) (*models.Department, error)
	GetAllDepartments(ctx context.Context) ([]*models.Department, error)
	GetDepartment(ctx context.Context, id int) (*models.Department, error)
	UpdateDepartment(ctx context.Context, dto *models.Department) (*models.Department, error)
	DeleteDepartment(ctx context.Context, id int) (bool, error)

	CheckDepartmentExists(ctx context.Context, department *models.Department) (bool, error)
}

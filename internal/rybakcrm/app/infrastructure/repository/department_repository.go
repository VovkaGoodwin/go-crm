package repository

import (
	"context"
	"crm-backend/internal/rybakcrm/app/domain/models"
	"github.com/jmoiron/sqlx"
)

type DepartmentRepository struct {
	db *sqlx.DB
}

func NewDepartmentRepositoryRepository(db *sqlx.DB) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (d *DepartmentRepository) CreateDepartment(ctx context.Context, department *models.Department) (*models.Department, error) {

	row := d.db.QueryRowxContext(
		ctx,
		"INSERT INTO departments (name, description) VALUES ($1, $2) RETURNING *;",
		department.Name,
		department.Description,
	)

	var result *models.Department

	if err := row.StructScan(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *DepartmentRepository) GetAllDepartments(ctx context.Context) ([]*models.Department, error) {

	var result []*models.Department
	err := d.db.SelectContext(
		ctx,
		&result,
		"SELECT id, name, description FROM departments",
	)

	return result, err
}

func (d *DepartmentRepository) GetDepartment(ctx context.Context, id int) (*models.Department, error) {
	var result models.Department

	err := d.db.GetContext(ctx, &result, "SELECT * FROM departments WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *DepartmentRepository) UpdateDepartment(ctx context.Context, department *models.Department) (*models.Department, error) {
	row := d.db.QueryRowxContext(
		ctx,
		"UPDATE departments SET name=$1, description=$2 WHERE id=$3 RETURNING *;",
		department.Name,
		department.Description,
		department.Id,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var result models.Department
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *DepartmentRepository) CheckDepartmentExists(ctx context.Context, department *models.Department) (bool, error) {
	var result bool

	err := d.db.GetContext(
		ctx,
		&result,
		"SELECT EXISTS (SELECT * FROM departments WHERE id <> $1 AND LOWER(name) = LOWER($2))",
		department.Id,
		department.Name,
	)

	if err != nil {
		return true, err
	}

	return result, nil
}

func (d *DepartmentRepository) DeleteDepartment(ctx context.Context, id int) (bool, error) {
	_, err := d.db.QueryxContext(
		ctx,
		"DELETE FROM departments WHERE id=$1",
		id,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

package employee

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindActiveSmithEmployees(ctx context.Context) ([]EmployeeName, error)
	FindEmployeesWithoutReviews(ctx context.Context) ([]EmployeeName, error)
	GetHireDateDiffActiveEmployees(ctx context.Context) (int, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) FindActiveSmithEmployees(ctx context.Context) ([]EmployeeName, error) {
	query := `SELECT first_name, last_name FROM employees WHERE termination_date IS NULL AND last_name ILIKE 'Smith%' ORDER BY last_name, first_name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []EmployeeName
	for rows.Next() {
		var e EmployeeName
		err := rows.Scan(&e.FirstName, &e.LastName)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func (r *repository) FindEmployeesWithoutReviews(ctx context.Context) ([]EmployeeName, error) {
	query := `SELECT e.first_name, e.last_name FROM employees e WHERE NOT EXIST (
				SELECT 1 FROM annual_reviews ar WHERE ar.employee_id = e.id 
			) ORDER BY e.hire_date`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []EmployeeName
	for rows.Next() {
		var e EmployeeName
		err := rows.Scan(&e.FirstName, &e.LastName)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func (r *repository) GetHireDateDiffActiveEmployees(ctx context.Context) (int, error) {
	query := `SELECT COALESCE(MAX(hire_date) - MIN(hire_date), 0) AS diff_days FROM employees WHERE termination_date IS NULL`

	var diffDays int
	err := r.db.QueryRow(ctx, query).Scan(&diffDays)
	if err != nil {
		return 0, nil
	}

	return diffDays, nil
}

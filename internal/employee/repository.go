package employee

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindActiveSmithEmployees(ctx context.Context) ([]EmployeeName, error)
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
